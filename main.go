package main

import (
	"minitok/config"
	"minitok/log"
	"minitok/routes"
	"minitok/storage"
	"minitok/usal"

	"github.com/gin-gonic/gin"
)

func main() {
	defer usal.CloseDataBase()
	defer usal.CloseRedis()
	defer log.Sync()
	//初始化项目配置
	config.LoadConfig() //加载配置信息
	log.InitLog()
	usal.InitDatabase()
	storage.InitMinio()
	usal.RedisInit()

	rou := gin.Default()
	rou = routes.SetRoute(rou)
	rou.Run()
}
