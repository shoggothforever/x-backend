package responses

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Id       string `form:"id" json:"id" binding:"required"`
	Password string `form:"psw" json:"psw" binding:"required"`
}
type Responses struct {
	Code    string
	Message string
	Data    string
}
type Claims struct {
	Id  string `json:"id"`
	Psw string `json:"psw"`
	jwt.StandardClaims
}

var AuthClaims *Claims
var AuthToken string

//Response data
func SendResponse(c *gin.Context, code int, msg string, data ...interface{}) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
