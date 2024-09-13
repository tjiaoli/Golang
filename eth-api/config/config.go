package config

import (
	"github.com/spf13/viper"
	"log"
)

func InitC() {
	viper.SetConfigName("config") // 配置文件名（不带扩展名）
	viper.AddConfigPath(".")      // 配置文件路径
	viper.SetConfigType("yaml")   // 配置文件类型

	err := viper.ReadInConfig() // 读取配置文件
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func GetInfuraUrl() string {
	return viper.GetString("infura.url")
}

func GetRedisAddr() string {
	return viper.GetString("redis.addr")
}

func GetMysqlDSN() string {
	return viper.GetString("mysql.dsn")
}
