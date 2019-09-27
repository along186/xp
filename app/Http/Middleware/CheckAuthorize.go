package Middleware

import (
	"github.com/gin-gonic/gin"
	"xp/app/Constant"
	"xp/pkg/Respone"
)

func CheckAuthorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		var msgCode int

		token := c.GetHeader("token")
		if token == "" {
			msgCode = Constant.IllegalRequest
		}

		if msgCode != 0 {
			Respone.SetContext(c).Error(Constant.GetMsg(msgCode))
			c.Abort()
			return
		}

		c.Next()
	}
}
