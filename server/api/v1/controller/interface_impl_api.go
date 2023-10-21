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
	data, total, err := interfaceImplService.FindInterfaceImplements(limit, offset, params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(response.PageResult{
		List:  data,
		Total: total,
	}, c)
}

// 根据id查询接口实现
func (iia *InterfaceImplApi) FindInterfaceImplById(c *gin.Context) {
	// 获取id
	id := c.Param("id")
	data, err := interfaceImplService.FindInterfaceImplById(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(data, c)
}

// 根据id更新接口实现
func (iia *InterfaceImplApi) UpdateInterfaceImplById(c *gin.Context) {
	// 获取id
	id := c.Param("id")
	var ii entity.InterfaceImpl
	if err := c.ShouldBindJSON(&ii); err != nil {
		response.FailWithDetail(nil, err.Error(), c)
		return
	}
	if err := interfaceImplService.UpdateInterfaceImplById(id, ii); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OK(c)
}

// 根据id删除接口实现
func (iia *InterfaceImplApi) DeleteInterfaceImplById(c *gin.Context) {
	// 获取id
	id := c.Param("id")
	if err := interfaceImplService.DeleteInterfaceImplById(id); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OK(c)
}

// 根据id切换接口实现状态
func (iia *InterfaceImplApi) SwitchInterfaceImplById(c *gin.Context) {
	var (
		ii entity.InterfaceImpl
		id string
	)
	id = c.Param("id")
	if err := c.ShouldBindJSON(&ii); err != nil {
		response.FailWithDetail(nil, err.Error(), c)
		return
	}
	if err := interfaceImplService.SwitchInterfaceImplById(id, ii.Enabled); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OK(c)
}
