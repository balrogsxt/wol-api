package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
	"strings"
	"wolapi/api"
	"wolapi/app"
)

var WindowsController = new(_WindowsController)

type _WindowsController struct {
}

//CallWol 调用网络唤醒操作
func (_WindowsController) CallWol(c *gin.Context) {
	_isForce := c.PostForm("isForce") //是否强制开机,不走检测
	isJumpIpCheck := false
	if strings.ToLower(_isForce) == "true" {
		isJumpIpCheck = true
	}
	conf := app.Config.Wol
	err := api.Wol.WakeOnLan(conf.MacAddress, conf.Network, conf.Ip, isJumpIpCheck)
	if err != nil {
		app.Response.FailMsg(c, err.Error())
		return
	}
	app.Response.Ok(c, gin.H{
		"ip":          conf.Ip,
		"mac_address": conf.MacAddress,
		"network":     conf.Network,
	}, "正在唤醒目标机器中")
}

//CallSleep 进行休眠操作
func (_WindowsController) CallSleep(c *gin.Context) {
	conf := app.Config.Wol
	status, err := api.Wol.IsPowerOn(conf.Ip)
	if err != nil {
		app.Response.FailMsg(c, err.Error())
		return
	}
	if !status {
		app.Response.FailMsg(c, "当前电脑已经处于休眠状态了")
		return
	}
	//内网地址
	callUrl := fmt.Sprintf("http://%s:%d/v1/windows/sleep", conf.Ip, conf.Port)

	if _, err := req.Post(callUrl); err != nil {
		//一般走这里,说明电脑要么是没开机,要么是网络错误
		app.Logger.Error("请求休眠操作失败: %s", err.Error())
		app.Response.FailMsg(c, "请求休眠操作失败,网络错误")
		return
	}
	//启动休眠操作
	app.Response.Ok(c, nil, "已提交休眠请求,等待目标电脑响应")
}

//PowerStatus 目标唤醒机器的开机状态
func (_WindowsController) PowerStatus(c *gin.Context) {
	conf := app.Config.Wol
	status, err := api.Wol.IsPowerOn(conf.Ip)
	if err != nil {
		app.Response.FailMsg(c, err.Error())
		return
	}
	app.Response.Ok(c, gin.H{
		"is_power": status,
	})
}
