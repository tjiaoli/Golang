package blockchain

import (
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
