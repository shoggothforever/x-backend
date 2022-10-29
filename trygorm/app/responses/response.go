package responses

import "github.com/gin-gonic/gin"

type Responses struct {
	Code    string
	Message string
	Data    string
}

//Response data
func SendResponse(c *gin.Context, code int, msg string, data ...interface{}) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
