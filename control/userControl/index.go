package usercontr

import (
	"go_project/structs/userstructs"

	"github.com/gin-gonic/gin"
)

// 登陆
func Login(ctx *gin.Context) {

}

// 注册
func Register(ctx *gin.Context) {
	req := new(userstructs.Register)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(422, map[string]interface{}{
			"errCode": -1,
			"msg":     "请求参数错误",
			"format":  1,
		})
	}
}
