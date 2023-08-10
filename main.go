package main

import (
	"TikTokLite/common"
	"TikTokLite/config"
	"TikTokLite/log"
	"TikTokLite/minioStore"
	"TikTokLite/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	MinitokInit() //初始化项目
	defer common.CloseDataBase()
	defer common.CloseRedis()
	defer log.Sync()

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
