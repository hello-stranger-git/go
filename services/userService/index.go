package userservice

import (
	usermodel "go_project/model/userModel"
	requestuserstructs "go_project/structs/request/requestUserStructs"

	"github.com/gin-gonic/gin"
)

type userServiceClass struct {
	ctx *gin.Context
}

func UserServiceInit(ctx *gin.Context) *userServiceClass {
	return &userServiceClass{ctx: ctx}
}
func (userServiceClass *userServiceClass) Register(req requestuserstructs.Register) error {
	userModelConnect := usermodel.GetUserModel()
	_, err := userModelConnect.Register(&usermodel.UserInfos{
		UserName: req.UserName,
		PassWord: req.PassWord,
	})
	if err != nil {
		return err
	}
	return nil
}
