package main

import (
	"minitok/common"
	"minitok/config"
	"minitok/log"
	"minitok/minioStore"
	"minitok/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	defer common.CloseDataBase()
	defer common.CloseRedis()
	defer log.Sync()

	MinitokInit() //初始化项目
	rou := gin.Default()
	rou = routes.SetRoute(rou)
	rou.Run()
}

func MinitokInit() {
	config.LoadConfig() //加载
	log.InitLog()
	common.InitDatabase()
	minioStore.InitMinio()
	common.RedisInit()
}
