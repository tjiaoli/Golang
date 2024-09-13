package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterEthRoutes() {
	r := gin.Default()
	r.GET("/eth/block/:block_num", GetBlock)
	r.GET("/eth/block/:block_num", GetTransaction)
	r.GET("/eth/block/:block_num", GetTransactionReceipt)
	// 启动服务器
	r.Run(":8080")
}
