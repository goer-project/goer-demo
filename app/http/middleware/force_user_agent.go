package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/goer-project/goer/response"
)

func ForceUserAgent() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Request.Header["User-Agent"]) == 0 {
			response.BadRequest(c, errors.New("User-Agent not found"))
			return
		}

		c.Next()
	}
}
