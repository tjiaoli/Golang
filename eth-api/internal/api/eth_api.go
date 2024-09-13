package api

import (
	"eth-api/internal/database"
	"eth-api/internal/database/dataRepository"
	"eth-api/internal/models"
	"eth-api/internal/service"
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
				block, blockTx, err := service.GetBlockFromChain(blockNum, full)
				if err != nil {
					//没有获取到区块信息
					c.JSON(http.StatusNotFound, gin.H{"blockDataError": err})
					return
				} else {
					//获取到区块信息存储到redis和数据库
					_, msg, state := service.SaveBlocks(blockNum, block, blockTx, full, true)
					c.JSON(state, gin.H{"message": msg, "block": block})
					return
				}
			}
		}
	} else {
		//查数据库
		blockData = blockRepo.GetBlockFromMySQL(blockNumInt)
		if blockData == nil {
			//从区块链中取数
			block, blockTx, err := service.GetBlockByNumber(blockNumInt, full)
			if err != nil {
				//没有获取到区块信息
				c.JSON(http.StatusNotFound, gin.H{"blockDataError": err})
				return
			} else {
				//获取到区块信息存储到redis和数据库
				_, msg, state := service.SaveBlocks(blockNum, block, blockTx, full, false)
				c.JSON(state, gin.H{"message": msg, "block": block})
				return
			}
		}
	}
	c.JSON(200, gin.H{"block": blockData})
}

// 获取交易数据 API
func GetTransaction(c *gin.Context) {
	txHash := c.Param("tx_hash")
	blockRepo := dataRepository.NewBlocksRepository()
	txData, err := blockRepo.GetTxFromMySQL(txHash)
	if err != nil {
		txData, err := service.GetTxFromChainByTxHash(txHash)
		if err != nil {
			var txList []*models.Transaction
			txList = append(txList, txData)
			err := blockRepo.SaveTxToMySQL(txList)
			if err != nil {
				c.JSON(400, gin.H{"Failed to save transaction to MySQL:": err})
			}
		}
		c.JSON(200, gin.H{"transaction": txData})
	} else {
		c.JSON(200, gin.H{"transaction": txData})
	}

}

// 获取交易存根数据 API
func GetTransactionReceipt(c *gin.Context) {
	txHash := c.Param("tx_hash")
	blockRepo := dataRepository.NewBlocksRepository()
	receiptData := blockRepo.GetReceiptFromMySQL(txHash)
	if receiptData == nil {
		receiptData, err := service.GetTransactionReceipt(txHash)
		if err == nil {
			err := blockRepo.SaveReceiptToMySQL(receiptData)
			if err != nil {
				c.JSON(200, gin.H{"receipt": receiptData})
			} else {
				c.JSON(400, gin.H{"error": err})
			}
		} else {
			c.JSON(400, gin.H{"error": err})
		}
	} else {
		c.JSON(200, gin.H{"receipt": receiptData})
	}

}
