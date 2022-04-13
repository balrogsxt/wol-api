package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"runtime/debug"
	"wolapi/app"
)
// Recover 异常捕捉
func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				app.Logger.Error("运行发生了一个未知错误: Url -> %s — Error -> %#v", c.Request.RequestURI, r)
				app.Logger.Error("\n%s", string(debug.Stack()))
				c.JSON(500,gin.H{
					"status":    1,
					"msg":       "欸,发生了一个未知的错误,阁下请稍后再试~",
				})
			}

		}()
		c.Next()
	}
}

// Cors 跨域请求
func Cors() gin.HandlerFunc {
	conf := cors.DefaultConfig()
	conf.AllowMethods = []string{"*"}
	conf.AllowHeaders = []string{"*,Content-Type,Authorization"}
	conf.AllowOrigins = []string{"*"}
	conf.AllowCredentials = true
	return cors.New(conf)
}
