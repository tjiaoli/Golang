package blockchain

import (
	"eth-api/internal/database"
	"eth-api/internal/database/dataRepository"
	"eth-api/internal/models"
	"fmt"
	"log"
)

// 从链上获取区块数据
func GetBlockFromChain(blockNum string, full bool) (*models.Block, []*models.Transaction, error) {
	if EthClient == nil {
		log.Println("Ethereum client is not initialized")
		return nil, nil, fmt.Errorf("Ethereum client is not initialized")
	}
	if blockNum == "head" {
		blockNum = "latest"
	}
	block, err := GetBlockByNumber(blockNum, full)

	if err != nil {
		log.Printf("Failed to get block from chain: %v", err)
		return nil, nil, err
	}

	if full {
		blockTx, blockTxerr := GetTxFromChain(blockNum)
		if blockTxerr != nil {
			return block, nil, blockTxerr
		} else {
			return block, blockTx, nil
		}
	}

	return block, nil, nil
}

func SaveBlocks(blockNum string, blockData *models.Block, blockTx []*models.Transaction, full bool) (*models.Block, string, int) {

	blockRepo := dataRepository.NewBlocksRepository()

	RedisErr := database.SaveBlockToRedis(blockNum, blockData)
	if RedisErr != nil {
		MysqlErr := blockRepo.SaveBlockToMySQL(blockData)
		if MysqlErr != nil {
			return blockData, "SaveBlockToRedisError: " + RedisErr.Error() + ", saveBlockToMySQLError: " + MysqlErr.Error(), 500
		} else {
			return blockData, RedisErr.Error(), 200
		}
	} else {
		MysqlErr := blockRepo.SaveBlockToMySQL(blockData)
		if MysqlErr != nil {
			return blockData, "saveBlockToMySQLError: " + MysqlErr.Error(), 500
		}
	}
	//如果full==true那需要存储交易数据
	if full {
		blockTxErr := blockRepo.SaveTxToMySQL(blockTx)
		if blockTxErr != nil {
			return blockData, "saveBlockTxToMySQLError:" + blockTxErr.Error(), 500
		}
	}
	return blockData, "查询并完成存储", 200
}
