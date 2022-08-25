package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.New()

	Router.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "未找到当前路由")
	})
	baseRouter := Router.Group("/api")
	BaseRouter(baseRouter)
	return Router
}
