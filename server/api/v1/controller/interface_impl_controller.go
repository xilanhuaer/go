package controller

import (
	"interface/model/common/response"
	"interface/model/entity"
	"interface/utils"

	"github.com/gin-gonic/gin"
)

type InterfaceImplController struct {
}

func (interfaceImplController *InterfaceImplController) Create(context *gin.Context) {
	var interfaceImpl entity.InterfaceImpl
	name := context.MustGet("username").(string)
	if err := context.ShouldBindJSON(&interfaceImpl); err != nil {
		response.FailWithDetail(nil, err.Error(), context)
		return
	}
	interfaceImpl.Creator = name
	interfaceImpl.Updator = name
	if err := interfaceImplService.Create(interfaceImpl); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OK(context)
}
func (interfaceImplController *InterfaceImplController) List(context *gin.Context) {
	limit, offset, err := utils.PageUtil(context.DefaultQuery("page", "1"), context.DefaultQuery("page_size", "10"))
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	params := map[string]string{
		"name":               context.DefaultQuery("name", ""),
		"type":               context.DefaultQuery("type", ""),
		"interface_name":     context.DefaultQuery("interface_name", ""),
		"main_collection_id": context.DefaultQuery("main_collection_id", ""),
		"sub_collection_id":  context.DefaultQuery("sub_collection_id", ""),
		"enabled":            context.DefaultQuery("enabled", ""),
	}
	data, total, err := interfaceImplService.List(limit, offset, params)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(response.PageResult{
		List:  data,
		Total: total,
	}, context)
}

// 根据id查询接口实现
func (interfaceImplController *InterfaceImplController) Find(context *gin.Context) {
	// 获取id
	id := context.Param("id")
	data, err := interfaceImplService.Find(id)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(data, context)
}

// 根据id更新接口实现
func (interfaceImplController *InterfaceImplController) Update(context *gin.Context) {
	// 获取id
	id := context.Param("id")
	name := context.MustGet("username").(string)
	var interfaceImpl entity.InterfaceImpl
	if err := context.ShouldBindJSON(&interfaceImpl); err != nil {
		response.FailWithDetail(nil, err.Error(), context)
		return
	}
	if err := interfaceImplService.Update(id, interfaceImpl, name); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OK(context)
}

// 根据id删除接口实现
func (interfaceImplController *InterfaceImplController) Delete(context *gin.Context) {
	// 获取id
	id := context.Param("id")
	name := context.MustGet("username").(string)
	if err := interfaceImplService.Delete(id, name); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OK(context)
}

// 根据id切换接口实现状态
func (interfaceImplController *InterfaceImplController) Enable(context *gin.Context) {
	var (
		interfaceImpl entity.InterfaceImpl
		id            string
		name          = context.MustGet("username").(string)
	)
	id = context.Param("id")
	if err := context.ShouldBindJSON(&interfaceImpl); err != nil {
		response.FailWithDetail(nil, err.Error(), context)
		return
	}
	interfaceImpl.Updator = name
	if err := interfaceImplService.Enable(id, interfaceImpl); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OK(context)
}
