package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"wolapi/app"
	"wolapi/http_service/routers"
	"time"
)

func main()  {
	rand.Seed(time.Now().UnixNano())
	//加载基础配置文件、日志库等
	if err := app.InitConfig();err != nil {
		log.Fatalln(fmt.Sprintf("初始化配置失败: %s",err.Error()))
	}
	if err := app.InitLogger();err != nil {
		log.Fatalln(fmt.Sprintf("初始化日志失败: %s",err.Error()))
	}
	app.Logger.Info("系统初始化成功")

	//启动http服务器
	http := gin.New()
	routers.RegisterRouter(http)

	if err := http.Run(fmt.Sprintf("%s:%d",app.Config.Http.Host,app.Config.Http.Port));err != nil {
		app.Logger.Fatal("启动Http服务器失败: %s",err.Error())
	}
}
