package routes

import (
	"minitok/common"
	"minitok/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	//函数创建一个名为 user 的路由组
	user := r.Group("user")
	{

		user.POST("/login/", controller.UserLogin)
		user.GET("/", common.AuthMiddleware(), controller.GetUserInfo)

		user.POST("/register/", controller.UserRegister)
	}

}
