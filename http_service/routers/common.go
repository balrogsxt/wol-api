package routers

import (
	"github.com/gin-gonic/gin"
	"wolapi/http_service/middlewares"
)

func RegisterRouter(g *gin.Engine) {
	//注册中间件
	g.Use(
		middlewares.Recover(),
		middlewares.Cors(),
		middlewares.Auth(),
	)
	registerApi(g)
}
