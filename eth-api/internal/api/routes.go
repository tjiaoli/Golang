package api

import (
	"context"
	"eth-api/internal/scheduler"
	"github.com/gin-gonic/gin"
)

func RegisterEthRoutes() {
	// 启动定时任务
	ctx, cancel := context.WithCancel(context.Background())
	scheduler.CancelFunc = cancel
	scheduler.Ctx = ctx
	go scheduler.StartScheduler()
	r := gin.Default()
	r.GET("/eth/block/:block_num", GetBlock)
	r.GET("/eth/block/:block_num", GetTransaction)
	r.GET("/eth/block/:block_num", GetTransactionReceipt)
	// 定义停止定时任务的接口
	r.GET("/stopScheduler", scheduler.StopScheduler)

	//启动定时任务
	r.GET("/startScheduler", func(c *gin.Context) {
		scheduler.StartScheduler() // 启动定时任务
		c.JSON(200, gin.H{"status": "定时任务已启动"})
	})
	//查看定时任务状态
	r.GET("/GetSchedulerStatus", scheduler.GetSchedulerStatus)
	// 启动服务器
	r.Run(":8080")
}
