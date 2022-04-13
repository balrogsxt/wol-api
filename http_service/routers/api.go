package routers

import (
	"github.com/gin-gonic/gin"
	"wolapi/http_service/controllers"
)

func registerApi(g *gin.Engine) {

	_wol := g.Group("/windows")
	{
		_wol.POST("/wol_call", controllers.WindowsController.CallWol)         //调用网络唤醒
		_wol.POST("/sleep_call", controllers.WindowsController.CallSleep)     //系统休眠
		_wol.POST("/power_status", controllers.WindowsController.PowerStatus) //获取目标机器状态
	}

}
