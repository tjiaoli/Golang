package scheduler

import (
	"context"
	"eth-api/internal/database"
	"eth-api/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var CancelFunc context.CancelFunc

var Ctx context.Context

/**停止定时任务*/
func StopScheduler(c *gin.Context) {
	if CancelFunc != nil {
		CancelFunc() // 调用 cancel 停止任务
		c.JSON(200, gin.H{"status": "定时任务已停止"})
	} else {
		c.JSON(400, gin.H{"status": "定时任务未启动"})
	}
}

var schedulerRunning bool

/**启动定时任务*/
func StartScheduler() {
	ticker := time.NewTicker(1 * time.Minute)
	schedulerRunning = true
	defer func() { schedulerRunning = false }()

	for {
		select {
		case <-ticker.C:
			// 并发获取区块并保存
			errs := make(chan error, 3)
			go retrieveAndSaveBlock("latest", errs)
			go retrieveAndSaveBlock("finalized", errs)
			go retrieveAndSaveBlock("safe", errs)

			// 收集并记录错误
			for i := 0; i < 3; i++ {
				if err := <-errs; err != nil {
					log.Printf("Error: %v", err)
				}
			}
		case <-Ctx.Done():
			log.Println("定时任务已停止")
			ticker.Stop()
			return
		}
	}
}

/**获取定时任务状态*/
func GetSchedulerStatus(c *gin.Context) {
	if schedulerRunning {
		c.JSON(200, gin.H{"status": "Scheduler is running"})
	} else {
		c.JSON(200, gin.H{"status": "Scheduler is stopped"})
	}
}

func retrieveAndSaveBlock(blockType string, errs chan<- error) {
	block, _, err := service.GetBlockFromChain(blockType, false)
	if err != nil {
		log.Printf("Failed to get head block: %v", err)
		errs <- err
		return
	} else {
		err := database.SaveBlockToRedis(blockType, block)
		if err != nil {
			log.Printf("Failed to save head block to redis: %v", err)

		}
	}
	errs <- err
}
