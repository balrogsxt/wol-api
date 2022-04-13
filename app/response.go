package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)
var Response = new(ApiResponse)
type ApiResponse struct {

}
//Ok 返回成功数据params=[1=data,2=msg,3=status]
func (a *ApiResponse) Ok(c *gin.Context,params ...interface{})  {
	var data interface{}
	var msg string = "終わりました"
	var status int64 = 200
	if len(params) >= 1 {
		data = params[0]
	}
	if len(params) >= 2 {
		msg = fmt.Sprintf("%v",params[1])
	}
	if len(params) >= 3 {
		status,_ = strconv.ParseInt(fmt.Sprintf("%v",params[2]),10,64)
	}
	a.render(c,msg,data,status,200)
}
func (a *ApiResponse) Fail(c *gin.Context,params ...interface{})  {
	var data interface{}
	var msg string = "System Fail"
	var status int64 = 500
	if len(params) >= 1 {
		data = params[0]
	}
	if len(params) >= 2 {
		msg = fmt.Sprintf("%v",params[1])
	}
	if len(params) >= 3 {
		status,_ = strconv.ParseInt(fmt.Sprintf("%v",params[2]),10,64)
	}
	a.render(c,msg,data,status,500)
}
func (a *ApiResponse) ParamsFail(c *gin.Context) {
	a.render(c,"参数错误",nil,1,500)
}
func (a *ApiResponse) FailMsg(c *gin.Context,msg string) {
	a.render(c,msg,nil,1,500)
}
func (ApiResponse) render(c *gin.Context,msg string,data interface{},status int64,httpStatus int)  {
	c.JSON(httpStatus,Map{
		"status":status,
		"msg":msg,
		"data":data,
	})
	return
}
