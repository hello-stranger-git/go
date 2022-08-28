package middlewares

import (
	"github.com/dgrijalva/jwt-go"
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
	return token.SignedString(j.SigninKey)
}

// 获取signKey
func GetSignKey() string {
	return APP_KEY
}
