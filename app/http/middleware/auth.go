package middleware

import (
	"goer/pkg/auth"

	"github.com/gin-gonic/gin"
	"github.com/goer-project/goer/response"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		if auth.Id(c) == 0 {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		c.Next()
	}
}
