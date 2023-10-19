package controller

import (
	"interface/model/common/response"
	"interface/model/entity"
	"interface/utils"

	"github.com/gin-gonic/gin"
)

type InterfaceImplApi struct {
}

func (iia *InterfaceImplApi) CreateImpl(c *gin.Context) {
	var ii entity.InterfaceImpl
	if err := c.ShouldBindJSON(&ii); err != nil {
		response.FailWithDetail(nil, err.Error(), c)
		return
	}
	if err := interfaceImplService.CreateImpl(ii); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OK(c)
}
func (iia *InterfaceImplApi) FindInterfaceImplements(c *gin.Context) {
	limit, offset, err := utils.PageUtil(c.DefaultQuery("page", "1"), c.DefaultQuery("page_size", "10"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	params := map[string]string{
		"name":               c.DefaultQuery("name", ""),
		"type":               c.DefaultQuery("type", ""),
		"interface_id":       c.DefaultQuery("interface_id", ""),
		"main_collection_id": c.DefaultQuery("main_collection_id", ""),
		"sub_collection_id":  c.DefaultQuery("sub_collection_id", ""),
		"enabled":            c.DefaultQuery("enabled", ""),
	}
	data, err := interfaceImplService.FindInterfaceImplements(limit, offset, params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(data, c)
}
