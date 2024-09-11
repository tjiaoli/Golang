package api

import "eth-api/internal/database"

func main() {
	//初始化mysql
	database.InitMysql()
	//初始化redis
	database.InitRedis()

}
