package main

import (
	"eth-api/config"
	"eth-api/internal/api"
	"eth-api/internal/blockchain"
	"eth-api/internal/database"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	//初始化config
	go func() {
		defer wg.Done()
		config.InitC()
		log.Println("Config initialized")
	}()

	// 等待配置初始化完成
	wg.Wait()

	//初始化mysql
	database.InitMysql()
	//初始化redis
	database.InitRedis()
	blockchain.InitEthC()
	//路由
	api.RegisterEthRoutes()
}
