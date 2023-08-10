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

	MinitokInit() //初始化项目
	rou := gin.Default()
	rou = routes.SetRoute(rou)
	rou.Run()
}

func MinitokInit() {
	config.LoadConfig() //加载
	log.InitLog()
	usal.InitDatabase()
	storage.InitMinio()
	usal.RedisInit()
}
