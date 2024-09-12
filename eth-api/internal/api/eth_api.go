package api

import (
	"eth-api/internal/blockchain"
	"eth-api/internal/database"
	"eth-api/internal/database/dataRepository"
	"eth-api/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 获取区块数据 API
func GetBlock(c *gin.Context) {
	blockNum := c.Param("block_num")
	//full := c.DefaultQuery("full", "false") == "true"

	blockRepo := dataRepository.NewBlocksRepository()

	var blockData *models.Block

	//判断是否为数字
	blockNumInt, err := strconv.Atoi(blockNum)
	if err != nil {
		if blockNum == "head" || blockNum == "finalized" || blockNum == "safe" {
			//查redis
			blockData, err = database.GetBlockFromRedis(blockNum)
			//从区块链中取数
			block, err := blockchain.GetBlockFromChain(blockNum, true)
			if err != nil {
				fmt.Println("blockDataError:", err)
				c.JSON(400, gin.H{"blockDataError": err})
				return
			} else {
				c.JSON(200, gin.H{"block": block})
				return
			}

		}

	} else {
		//查数据库
		blockData = blockRepo.GetBlockFromMySQL(blockNumInt)
	}

	c.JSON(200, gin.H{"block": blockData})
}
