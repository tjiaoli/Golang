package service

import (
	"context"
	"encoding/json"
	"eth-api/internal/blockchain"
	"eth-api/internal/database"
	"eth-api/internal/database/dataRepository"
	"eth-api/internal/models"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"time"
)

// 从链上获取区块数据
func GetBlockFromChain(blockNum string, full bool) (*models.Block, []*models.Transaction, error) {
	if blockchain.EthClient == nil {
		log.Println("Ethereum client is not initialized")
		return nil, nil, fmt.Errorf("Ethereum client is not initialized")
	}
	if blockNum == "safe" {
		blockNum = "pending"
	}
	block, err := blockchain.GetBlockByNumber(blockNum, full)

	if err != nil {
		log.Printf("Failed to get block from chain: %v", err)
		return nil, nil, err
	}

	if full {
		blockTx, blockTxerr := blockchain.GetTxFromChain(blockNum)
		if blockTxerr != nil {
			return block, nil, blockTxerr
		} else {
			return block, blockTx, nil
		}
	}

	return block, nil, nil
}

func SaveBlocks(blockNum string, blockData *models.Block, blockTx []*models.Transaction, full bool, saveRedis bool) (*models.Block, string, int) {

	blockRepo := dataRepository.NewBlocksRepository()

	if saveRedis {
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

func GetBlockByNumber(blockNumInt int, full bool) (*models.Block, []*models.Transaction, error) {
	if blockchain.EthClient == nil {
		log.Println("Ethereum client is not initialized")
		return nil, nil, fmt.Errorf("Ethereum client is not initialized")
	}
	blockNumber := big.NewInt(int64(blockNumInt))
	block, err := blockchain.EthClient.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Printf("Failed to fetch block by number %d: %v", blockNumInt, err)
		return nil, nil, err
	}
	// 将 types.Block 转换为 models.Block
	blockData := &models.Block{
		BlockHash:        block.Hash().Hex(),
		BlockNumber:      block.Number().Int64(),
		ParentHash:       block.ParentHash().Hex(),
		Timestamp:        int64(block.Time()),
		Miner:            block.Coinbase().Hex(),
		GasUsed:          int64(block.GasUsed()),
		GasLimit:         int64(block.GasLimit()),
		Size:             int64(block.Size()),
		TransactionCount: len(block.Transactions()), // 获取交易数量
	}
	// 如果需要完整交易信息
	var transactions []*models.Transaction
	if full {
		for _, tx := range block.Transactions() {
			var fromAddress string
			chainID, err := blockchain.EthClient.NetworkID(context.Background())
			if err != nil {
				log.Fatal(err)
			}
			if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
				fmt.Println("sender", sender.Hex())
				fromAddress = sender.Hex()
			}

			// 创建 Transaction 结构体并填充交易信息
			transactionData := &models.Transaction{
				TxHash:      tx.Hash().Hex(),
				BlockHash:   block.Hash().Hex(),
				BlockNumber: block.Number().Int64(),
				FromAddress: fromAddress,
				ToAddress:   tx.To().Hex(),
				Value:       tx.Value(), // 如果需要精确处理，可能需要其他方法
				GasPrice:    tx.GasPrice().Int64(),
				GasUsed:     int64(tx.Gas()),
				Nonce:       int64(tx.Nonce()),
				InputData:   fmt.Sprintf("%x", tx.Data()),
			}
			transactions = append(transactions, transactionData)
		}
	}
	return blockData, transactions, nil
}

func GetTxFromChainByTxHash(txHash string) (*models.Transaction, error) {
	if blockchain.EthClient == nil {
		log.Println("Ethereum client is not initialized")
		return nil, fmt.Errorf("Ethereum client is not initialized")
	}
	blockHash := common.HexToHash(txHash)
	tx, isPending, err := blockchain.EthClient.TransactionByHash(context.Background(), blockHash)

	if isPending {
		fmt.Println("Transaction is pending")
	} else {
		fmt.Println("Transaction is confirmed")
	}

	if err != nil {
		//log.Printf("Failed to get transaction from chain: %v", err)
		return nil, fmt.Errorf("Failed to get transaction from chain: %v", err)
	}

	// 手动将返回的 map 数据解析到 Transaction 结构体
	var fromAddress string
	chainID, err := blockchain.EthClient.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
		fmt.Println("sender", sender.Hex())
		fromAddress = sender.Hex()
	}
	transaction := &models.Transaction{
		TxHash:    tx.Hash().Hex(),
		BlockHash: txHash,
		//BlockNumber: tx.
		FromAddress: fromAddress,
		ToAddress:   tx.To().Hex(),
		Value:       tx.Value(), // 如果需要精确处理，可能需要其他方法
		GasPrice:    tx.GasPrice().Int64(),
		GasUsed:     int64(tx.Gas()),
		Nonce:       int64(tx.Nonce()),
		InputData:   fmt.Sprintf("%x", tx.Data()),
	}
	return transaction, nil
}

func GetTransactionReceipt(txHash string) (*models.TransactionReceipt, error) {
	// 调用以太坊客户端获取交易收据
	blockHash := common.HexToHash(txHash)
	txReceipt, err := blockchain.EthClient.TransactionReceipt(context.Background(), blockHash)
	if err != nil {
		log.Printf("Failed to get transaction receipt from chain: %v", err)
		return nil, err
	}

	logsJSON, err := json.Marshal(txReceipt.Logs)
	if err != nil {
		log.Printf("Failed to marshal logs: %v", err)
		return nil, err
	}
	// 构建 TransactionReceipt 结构体
	transactionReceipt := &models.TransactionReceipt{
		TxHash:            txReceipt.TxHash.Hex(),          // 转换为字符串
		BlockHash:         txReceipt.BlockHash.Hex(),       // 转换为字符串
		BlockNumber:       txReceipt.BlockNumber,           // uint64
		CumulativeGasUsed: txReceipt.CumulativeGasUsed,     // uint64
		GasUsed:           txReceipt.GasUsed,               // uint64
		ContractAddress:   txReceipt.ContractAddress.Hex(), // string
		Status:            txReceipt.Status,                // uint64
		Logs:              string(logsJSON),                // 日志的JSON字符串
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	return transactionReceipt, nil
}
