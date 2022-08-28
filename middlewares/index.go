package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// jwt信息
type CustomClaims struct {
	ID       int    `json:"userId"`
	UserName string `json:"userName"`
	jwt.StandardClaims
}

const (
	APP_KEY = "www.blog.com"
)

// jwt签名结构
type JWT struct {
	SigninKey []byte
}

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// 生成token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.SigninKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

// 检测token中间件
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status": -1,
				"msg":    "请求未携带token，无权限访问",
			})
			ctx.Abort()
			return
		}
		token = strings.ReplaceAll(token, "Bearer ", "")
		// j := NewJWT()
		_, err := Alysistoken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"errCode": -1,
				"msg":     err.Error(),
				"format":  1,
			})
			ctx.Abort()
		}
		ctx.Next()
	}
}

// 解析token
func Alysistoken(token string) (*CustomClaims, error) {
	customClaims := new(CustomClaims)
	claims, err := jwt.ParseWithClaims(token, customClaims, func(t *jwt.Token) (interface{}, error) {
		return []byte(GetSignKey()), nil
	})
	if err != nil {
		return nil, errors.New("解析token错误")
	}
	if !claims.Valid {
		return nil, errors.New("无效token")
	}
	return customClaims, nil
}

// 获取signKey
func GetSignKey() string {
	return APP_KEY
}
