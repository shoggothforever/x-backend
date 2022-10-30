package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"trygorm/app/responses"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// Set default variable
		c.Set("id", 149)
		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

//1.5.3 auth middleware
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token := c.Param("token"); token != responses.AuthToken {
			c.Set("foo", "foo")
			c.JSON(200, gin.H{
				"code": 400,
			})
			c.AbortWithStatusJSON(200, gin.H{
				"code": 401,
			})
			return
		} else {
			c.Set("todo", "todo")
		}

		c.Next()
	}
}
