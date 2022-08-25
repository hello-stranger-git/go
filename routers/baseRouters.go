package routers

import (
	usercontrol "go_project/control/userControl"

	"github.com/gin-gonic/gin"
)

func BaseRouter(router *gin.RouterGroup) {
	//登录
	router.POST("/login", usercontrol.Login)
	//注册
	router.POST("/regist", usercontrol.Register)
}
