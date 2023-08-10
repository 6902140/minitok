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
	r := gin.Default()
	r = routes.SetRoute(r)
	r.Run()
}

func MinitokInit() {
	config.LoadConfig() //加载
	log.InitLog()
	common.InitDatabase()
	minioStore.InitMinio()
	common.RedisInit()
}
