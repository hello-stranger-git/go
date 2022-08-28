package usercontr

import (
	"fmt"
	userservice "go_project/services/userService"
	requestuserstructs "go_project/structs/request/requestUserStructs"

	"github.com/gin-gonic/gin"
)

// 登陆
func Login(ctx *gin.Context) {
	req := new(requestuserstructs.Login)
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.JSON(422, map[string]interface{}{
			"errCode": -1,
			"msg":     "请求参数错误",
			"format":  1,
		})
	}
	err = userservice.UserServiceInit(ctx).Login(*req, ctx)
	fmt.Print(err)
	if err != nil {
		ctx.JSON(422, map[string]interface{}{
			"errCode": -1,
			"msg":     err.Error(),
			"format":  1,
		})
		return
	}
}

// 注册
func Register(ctx *gin.Context) {
	req := new(requestuserstructs.Register)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(422, map[string]interface{}{
			"errCode": -1,
			"msg":     "请求参数错误",
			"format":  1,
		})
	}
	err = userservice.UserServiceInit(ctx).Register(*req)
	if err != nil {
		ctx.JSON(422, map[string]interface{}{
			"errCode": -1,
			"msg":     err.Error(),
			"format":  1,
		})
		return
	}
}

// 获取用户信息
func GetUserInfo(ctx *gin.Context) {
	fmt.Println("获取用户信息")
}
