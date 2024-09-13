package database

import (
	"context"
	"encoding/json"
	"eth-api/config"
	"eth-api/internal/models"
	"github.com/go-redis/redis/v8"
	"log"
)

var redisClient *redis.Client
var ctx = context.Background()

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: config.GetRedisAddr(),
		DB:   0, // 默认数据库
	})
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}

// 从Redis中获取区块数据
func GetBlockFromRedis(blockIdentifier string) (*models.Block, error) {
	blockData, err := redisClient.Get(ctx, blockIdentifier).Result()
	if err == redis.Nil {
		return nil, err
	} else if err != nil {
		log.Printf("Failed to get block from Redis: %v", err)
	}

	// 检查 blockData 是否为 "null" 或空字符串
	if blockData == "" || blockData == "null" {
		log.Printf("Block data is null or empty for identifier: %s", blockIdentifier)
		return nil, nil // 这里可以选择返回 nil 或者自定义的错误
	}

	var block models.Block // 假设这是你定义的区块结构体
	jsonerr := json.Unmarshal([]byte(blockData), &block)
	if jsonerr != nil {
		log.Printf("Failed to unmarshal block data: %v", err)
		return nil, jsonerr
	}

	return &block, err
}

// 将区块数据保存到Redis
func SaveBlockToRedis(blockIdentifier string, blockData *models.Block) error {
	return saveToRedis(blockIdentifier+"_Block", blockData)
}

// 保存交易数据
func SaveTxToRedis(blockIdentifier string, tx *models.Transaction) error {
	return saveToRedis(blockIdentifier+"_TX", tx)
}

func saveToRedis(blockIdentifier string, data interface{}) error {
	jsonData, jsonerr := json.Marshal(data)
	if jsonerr != nil {
		log.Printf("Failed to marshal tx data: %v", jsonerr)
		return jsonerr
	}

	err := redisClient.Set(ctx, blockIdentifier, jsonData, 0).Err()
	if err != nil {
		log.Printf("Failed to save tx to Redis: %v", err)
		return err
	}
	return nil
}
