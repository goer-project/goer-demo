package middleware

import (
	"goer/app/models/user"
	"goer/global"
	"goer/pkg/auth"

	"github.com/gin-gonic/gin"
)

func Guard() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Init user info to gin.context
		c.Set("auth.user_id", 0)
		c.Set("auth.user", user.User{})

		// Parse token
		claims, err := auth.NewJWT().ParseToken(c)
		if err != nil {
			return
		}

		if claims.Guard != "" {
			return
		}

		var userInfo user.User
		global.DB.First(&userInfo, claims.ID)
		if userInfo.ID == 0 {
			return
		}

		// Set user info to gin.context
		c.Set("auth.user_id", userInfo.ID)
		c.Set("auth.user", userInfo)

		c.Next()
	}
}
