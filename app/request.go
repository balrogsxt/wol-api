package app

import "github.com/gin-gonic/gin"

var Request = new(ApiRequest)

type ApiRequest struct {

}

func (ApiRequest) ParseJson(c *gin.Context,obj interface{}) error {
	return c.ShouldBindJSON(&obj)
}
func (ApiRequest) ParseXml(c *gin.Context,obj interface{}) error {
	return c.ShouldBindXML(&obj)
}