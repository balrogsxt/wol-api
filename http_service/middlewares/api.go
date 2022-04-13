package middlewares

import (
	"github.com/gin-gonic/gin"
	"wolapi/app"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != app.Config.ApiKey {
			c.Abort()
			app.Response.FailMsg(c,"身份令牌鉴权无效")
			return
		}
		c.Next()
	}
}