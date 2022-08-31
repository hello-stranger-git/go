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
	err = userservice.UserServiceInit(ctx).Login(*req)
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

//根据id查询用户
func GetUserInfoById(ctx *gin.Context) {
	req := new(requestuserstructs.UserInfo)
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.JSON(422, map[string]interface{}{
			"errCode": -1,
			"msg":     "请求参数错误",
			"format":  1,
		})
		return
	}
	data, err := userservice.UserServiceInit(ctx).GetUserInfoById(*req)
	if err != nil {

		ctx.JSON(422, map[string]interface{}{
			"errCode": -1,
			"msg":     err.Error(),
			"format":  1,
		})
		return
	}
	ctx.JSON(200, map[string]interface{}{
		"errCode": 0,
		"msg":     "获取成功",
		"data":    data,
		"format":  1,
	})
}
