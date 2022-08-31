package userservice

import (
	"errors"
	responsemodel "go_project/model/responseModel"
	usermodel "go_project/model/userModel"
	requestuserstructs "go_project/structs/request/requestUserStructs"
	responseuserstructs "go_project/structs/response/responseUserStructs"
	"go_project/utils"
	"go_project/utils/jwt"

	"github.com/gin-gonic/gin"
)

type userServiceClass struct {
	ctx *gin.Context
}

func UserServiceInit(ctx *gin.Context) *userServiceClass {
	return &userServiceClass{ctx: ctx}
}

// 注册
func (userServiceClass *userServiceClass) Register(req requestuserstructs.Register) error {
	userModelConnect := usermodel.GetUserModel()
	_, err := userModelConnect.Register(&usermodel.UserInfos{
		UserName: req.UserName,
		PassWord: utils.MD5(req.PassWord),
	})
	if err != nil {
		return err
	}
	return nil
}

// 登陆
func (userServiceClass *userServiceClass) Login(req requestuserstructs.Login) error {
	userModelConnect := usermodel.GetUserModel()
	userInfo, err := userModelConnect.Login(req.UserName)
	if err != nil {
		return errors.New("用户获取失败")
	}
	PassWord := utils.MD5(req.PassWord)
	if PassWord != userInfo.PassWord {
		return errors.New("密码或账号错误")
	}
	data := responseuserstructs.UserInfo{
		Id:       userInfo.Id,
		Phone:    userInfo.Phone,
		Username: userInfo.UserName,
	}
	jwt.CreateToken(userInfo, data, userServiceClass.ctx)
	return nil
}

//根据id查询用户
func (userServiceClass *userServiceClass) GetUserInfoById(req requestuserstructs.UserInfo) (*responsemodel.UserInfo, error) {
	userInfo, err := usermodel.GetUserModel().SetId(req.ID).GetUserInfoById()
	if err != nil {
		return nil, errors.New("用户信息查询失败")
	}
	return userInfo, nil
}
