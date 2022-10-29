package middleware

import "github.com/gin-gonic/gin"

func Auth(r *gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("id", 149)
		c.Next()
	}
}
