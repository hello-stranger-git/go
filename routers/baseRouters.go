package routers

import "github.com/gin-gonic/gin"

func BaseRouter(router *gin.RouterGroup) {
	//登录
	router.POST("login")
}
