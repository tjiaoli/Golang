package api

import (
	"eth-api/internal/blockchain"
	"eth-api/internal/database"
	"eth-api/internal/database/dataRepository"
	"eth-api/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 获取区块数据 API
func GetBlock(c *gin.Context) {
	blockNum := c.Param("block_num")
	full := c.DefaultQuery("full", "false") == "true"

	blockRepo := dataRepository.NewBlocksRepository()

	var blockData *models.Block

	//判断是否为数字
	blockNumInt, err := strconv.Atoi(blockNum)
	if err != nil {
		if blockNum == "head" || blockNum == "finalized" || blockNum == "safe" {
			//查redis
			blockData, err = database.GetBlockFromRedis(blockNum)
			if blockData == nil {
				//从区块链中取数
				block, err := blockchain.GetBlockFromChain(blockNum, full)
				if err != nil {
					//没有获取到区块信息
					c.JSON(http.StatusNotFound, gin.H{"blockDataError": err})
					return
				} else {
					//获取到区块信息存储到redis和数据库
					_, msg, state := blockchain.SaveBlocks(blockNum, block, full)
					c.JSON(state, gin.H{"message": msg, "block": block})
					return
				}
			}
		}

	} else {
		//查数据库
		blockData = blockRepo.GetBlockFromMySQL(blockNumInt)
	}

	c.JSON(200, gin.H{"block": blockData})
}
