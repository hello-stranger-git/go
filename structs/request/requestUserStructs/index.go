package requestuserstructs

import (
	"gorm.io/gorm"
)

// 接收前端请求的注册数据模型
type Register struct {
	gorm.Model
	UserName string `form:"username" json:"username" binding:"required"`
	PassWord string `form:"password" json:"password" binding:"required"`
}

// 接收前端请求的登陆数据模型
type Login struct {
	gorm.Model
	UserName string `form:"username" json:"username" binding:"required"`
	PassWord string `form:"password" json:"password" binding:"required"`
}
