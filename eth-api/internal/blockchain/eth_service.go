package blockchain

import (
	"eth-api/internal/database"
	"eth-api/internal/database/dataRepository"
	"eth-api/internal/models"
	"fmt"
	"log"
)

// 从链上获取区块数据
func GetBlockFromChain(blockNum string, full bool) (*models.Block, error) {
	if EthClient == nil {
		log.Println("Ethereum client is not initialized")
		return nil, fmt.Errorf("Ethereum client is not initialized")
	}
	if blockNum == "head" {
		blockNum = "latest"
	}
	block, err := GetBlockByNumber(blockNum, full)

	if err != nil {
		log.Printf("Failed to get block from chain: %v", err)
		return nil, err
	}
	return block, nil
}

func SaveBlocks(blockNum string, blockData *models.Block, full bool) (*models.Block, string, int) {

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
	return blockData, "查询并完成存储", 200
}
