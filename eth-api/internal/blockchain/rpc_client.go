package blockchain

import (
	"context"
	"eth-api/internal/models"
	"fmt"
	"log"
	"strconv"
	"time"
)

// 从链上获取区块数据
func GetBlockByNumber(blockIdentifier string, full bool) (*models.Block, error) {
	var rawBlockData map[string]interface{}
	err := RpcClient.CallContext(context.Background(), &rawBlockData, "eth_getBlockByNumber", blockIdentifier, full)

	if err != nil {
		log.Printf("Failed to get block from chain: %v", err)
		return nil, err
	}
	//var blockData models.Block
	var parseErrors []string // 用于记录解析错误
	var (
		blockNumber      int64
		timestamp        int64
		gasLimit         int64
		gasUsed          int64
		size             int64
		transactionCount int
	)
	// 安全地解析区块数据并进行类型断言
	blockHash, ok := rawBlockData["hash"].(string)
	if !ok {
		parseErrors = append(parseErrors, "Failed to parse blockHash")
	}

	blockNumberStr, ok := rawBlockData["number"].(string)
	if !ok {
		parseErrors = append(parseErrors, "Failed to parse blockNumber")
	} else {
		blockNumber = parseStringToInt64(blockNumberStr)
	}

	parentHash, ok := rawBlockData["parentHash"].(string)
	if !ok {
		parseErrors = append(parseErrors, "Failed to parse parentHash")
	}

	timestampStr, ok := rawBlockData["timestamp"].(string)
	if !ok {
		parseErrors = append(parseErrors, "Failed to parse timestamp")
	} else {
		timestamp = parseStringToInt64(timestampStr)
	}

	miner, ok := rawBlockData["miner"].(string)
	if !ok {
		parseErrors = append(parseErrors, "Failed to parse miner")
	}

	gasLimitStr, ok := rawBlockData["gasLimit"].(string)
	if !ok {
		parseErrors = append(parseErrors, "Failed to parse gasLimit")
	} else {
		gasLimit = parseStringToInt64(gasLimitStr)
	}

	gasUsedStr, ok := rawBlockData["gasUsed"].(string)
	if !ok {
		parseErrors = append(parseErrors, "Failed to parse gasUsed")
	} else {
		gasUsed = parseStringToInt64(gasUsedStr)
	}

	sizeStr, ok := rawBlockData["size"].(string)
	if !ok {
		parseErrors = append(parseErrors, "Failed to parse size")
	} else {
		size = parseStringToInt64(sizeStr)
	}

	if full {
		transactions, ok := rawBlockData["transactions"].([]interface{})
		if !ok {
			parseErrors = append(parseErrors, "Failed to parse transactions")
		} else {
			transactionCount = len(transactions)
		}
	} else {
		transactionCount = 0
	}

	// 检查是否存在关键字段解析失败
	if len(parseErrors) > 0 {
		log.Println("Errors occurred while parsing block data:", parseErrors)
		// 这里可以根据需要返回错误，或者只是记录日志并继续处理
		return nil, fmt.Errorf("failed to parse block data: %v", parseErrors)
	}

	// 构造 Block 结构体
	blockData := &models.Block{
		BlockHash:        blockHash,
		BlockNumber:      blockNumber,
		ParentHash:       parentHash,
		Timestamp:        timestamp,
		Miner:            miner,
		GasLimit:         gasLimit,
		GasUsed:          gasUsed,
		Size:             size,
		TransactionCount: transactionCount,
		CreatedAt:        time.Now(),
	}
	return blockData, nil
}

func parseStringToInt64(value string) int64 {
	if parsedValue, err := strconv.ParseInt(value, 0, 64); err == nil {
		return parsedValue
	}
	log.Printf("Failed to parse string to int64: %v", value) // Logging the error
	return 0                                                 // Return default value or handle error appropriately
}

func GetTxFromChain(blockIdentifier string) ([]*models.Transaction, error) {
	var txData map[string]interface{}
	//err := ethClient.Call(&txData, "eth_getTransactionByHash", txHash)
	err := RpcClient.CallContext(context.Background(), &txData, "eth_getBlockByNumber", blockIdentifier, true)

	if err != nil {
		log.Printf("Failed to get transaction from chain: %v", err)
		return nil, fmt.Errorf("Failed to get transaction from chain: %v", err)
	}

	//fmt.Println(txData)
	// 手动将返回的 map 数据解析到 Transaction 结构体
	transactions, ok := txData["transactions"].([]interface{})
	if ok {
		var txList []*models.Transaction
		for _, tx := range transactions {
			txData, ok := tx.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("transaction data is not in expected format")
			}
			transaction := &models.Transaction{
				TxHash:      safeString(txData, "hash"),
				BlockHash:   safeString(txData, "blockHash"),
				BlockNumber: parseStringToInt64(safeString(txData, "blockNumber")),
				FromAddress: safeString(txData, "from"),
				ToAddress:   safeString(txData, "to"),
				Value:       parseStringToDecimal(safeString(txData, "value")),
				GasPrice:    parseStringToInt64(safeString(txData, "gasPrice")),
				GasUsed:     parseStringToInt64(safeString(txData, "gas")),
				Nonce:       parseStringToInt64(safeString(txData, "nonce")),
				InputData:   safeString(txData, "input"),
				CreatedAt:   time.Now(),
			}
			txList = append(txList, transaction)
		}
		return txList, nil
	} else {
		return nil, fmt.Errorf("transaction data is not in expected format")
	}

}

func safeString(data map[string]interface{}, key string) string {
	if val, ok := data[key].(string); ok {
		return val
	}
	return ""
}

func parseStringToDecimal(value string) float64 {
	if parsedValue, err := strconv.ParseFloat(value, 64); err == nil {
		return parsedValue
	}
	log.Printf("Failed to parse string to decimal: %v", value) // Logging the error
	return 0.0                                                 // Return default value or handle error appropriately
}
