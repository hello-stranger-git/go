package jwt

import (
	"go_project/middlewares"
	usermodel "go_project/model/userModel"
	responseuserstructs "go_project/structs/response/responseUserStructs"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateToken(userInfo *usermodel.UserInfos, data responseuserstructs.UserInfo, ctx *gin.Context) {
	j := middlewares.NewJWT()
	claims := middlewares.CustomClaims{
		ID:       userInfo.Id,
		UserName: userInfo.UserName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix()),            // 签名生效时间
			IssuedAt:  int64(time.Now().Unix()),            // 签发时间
			ExpiresAt: int64(time.Now().Unix() + 12*60*60), // 过期时间 12小时
			Issuer:    middlewares.GetSignKey(),            // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
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
		"data": responseuserstructs.LoginResponse{
			Token:    token,
			UserInfo: data,
		},
		"msg":    "登陆成功",
		"format": 1,
	})
}
