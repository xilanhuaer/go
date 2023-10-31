package controller

import (
	"interface/model/common/response"
	"interface/model/entity"
	"interface/orm"
	"log"

	"github.com/gin-gonic/gin"
)

type InterfaceController struct {
}

func (interfaceController *InterfaceController) Create(context *gin.Context) {
	var (
		e        entity.Interface
		username = context.MustGet("username").(string)
	)
	err := context.ShouldBindJSON(&e)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	err = orm.Create(&e, username)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(true, context)
}

func (interfaceController *InterfaceController) List(context *gin.Context) {
	page := context.DefaultQuery("page", "1")
	page_size := context.DefaultQuery("page_size", "10")
	params := map[string]string{
		"name":               context.DefaultQuery("name", ""),
		"type":               context.DefaultQuery("type", ""),
		"enabled":            context.DefaultQuery("enabled", ""),
		"created_at":         context.DefaultQuery("created_at", ""),
		"main_collection_id": context.DefaultQuery("main_collection_id", ""),
		"sub_collection_id":  context.DefaultQuery("sub_collection_id", ""),
	}
	e, count, err := interfaceService.List(page, page_size, params)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	data := response.PageResult{
		List:  e,
		Total: count,
	}
	response.OKWithData(data, context)
}
func (interfaceController *InterfaceController) Find(context *gin.Context) {
	id := context.Param("id")
	data, err := interfaceService.Find(id)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(data, context)
}
func (interfaceController *InterfaceController) Update(context *gin.Context) {
	var (
		e    entity.Interface
		name = context.MustGet("username").(string)
	)
	if err := context.ShouldBindJSON(&e); err != nil {
		log.Println(1)
		response.FailWithMessage(err.Error(), context)
		return
	}
	err := orm.Update(&e, name)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(true, context)
}
func (interfaceController *InterfaceController) Enable(context *gin.Context) {
	id := context.Param("id")
	name := context.MustGet("username").(string)
	var e entity.Interface
	if err := context.ShouldBindJSON(&e); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	err := orm.Enable(&e, id, name)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(true, context)
}
func (interfaceController *InterfaceController) Delete(context *gin.Context) {
	id := context.Param("id")
	name := context.MustGet("username").(string)
	if err := orm.Delete(&entity.Interface{}, id, name); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(true, context)
}
