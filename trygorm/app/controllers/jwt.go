package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"time"
	"trygorm/app/responses"
)

func GenerateJwt(id, psw string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(4 * time.Hour)
	claims := responses.Claims{
		id,
		psw,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "dsm",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte("dsm"))
	return token, err
}

func ParseToken(token string) (*responses.Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &responses.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("dsm"), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*responses.Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
func Init() {
	token, _ := GenerateJwt("todo", "todo")
	responses.AuthClaims, _ = ParseToken(token)
}
func Login(c *gin.Context) {
	var login responses.Login
	if err := c.ShouldBind(&login); err != nil {
		logrus.Error(err)
	}
	responses.AuthToken, _ = GenerateJwt(login.Id, login.Password)
	responses.AuthClaims, _ = ParseToken(responses.AuthToken)
	fmt.Println(responses.AuthClaims)
	if login.Id == "todo" && login.Password == "todo" {
		c.JSON(200, gin.H{
			"code":  200,
			"msg":   "登录成功",
			"token": responses.AuthToken,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "登录失败",
		})
	}
}
func Checkerr(err error) {
	if err != nil {
		log.Fatalln(err)
		return
	}
}
