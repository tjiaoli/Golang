package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var redisClient *redis.Client
var ctx = context.Background()

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "192.168.190.157:6379",
		DB:   0, // 默认数据库
	})
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}
