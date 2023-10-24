package controller

import (
	"interface/model/common/response"
	"interface/model/entity"
	"log"

	"github.com/gin-gonic/gin"
)

type InterfaceApi struct {
}

func (i *InterfaceApi) CreateInterface(c *gin.Context) {
	var e entity.Interface
	name := c.MustGet("userName").(string)
	err := c.ShouldBindJSON(&e)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	e.Creator = name
	e.Updator = name
	err = interfaceService.CreateInterface(e)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(true, c)
}

func (i *InterfaceApi) FindInterfaces(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	page_size := c.DefaultQuery("page_size", "10")
	params := map[string]string{
		"name":               c.DefaultQuery("name", ""),
		"type":               c.DefaultQuery("type", ""),
		"enabled":            c.DefaultQuery("enabled", ""),
		"created_at":         c.DefaultQuery("created_at", ""),
		"main_collection_id": c.DefaultQuery("main_collection_id", ""),
		"sub_collection_id":  c.DefaultQuery("sub_collection_id", ""),
	}
	e, count, err := interfaceService.FindInterfaces(page, page_size, params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data := response.PageResult{
		List:  e,
		Total: count,
	}
	response.OKWithData(data, c)
}
func (i *InterfaceApi) FindInterface(c *gin.Context) {
	id := c.Param("id")
	data, err := interfaceService.FindInterface(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(data, c)
}
func (i *InterfaceApi) UpdateInterface(c *gin.Context) {
	id := c.Param("id")
	name := c.MustGet("userName").(string)
	var e entity.Interface
	e.Updator = name
	if err := c.ShouldBindJSON(&e); err != nil {
		log.Println(1)
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := interfaceService.UpdateInterface(id, e)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(true, c)
}
func (i *InterfaceApi) CheckInterfaceEnable(c *gin.Context) {
	id := c.Param("id")
	name := c.MustGet("userName").(string)
	var e entity.Interface
	if err := c.ShouldBindJSON(&e); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	e.Updator = name
	err := interfaceService.CheckInterfaceEnable(id, e)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(true, c)
}
func (i *InterfaceApi) DeleteInterface(c *gin.Context) {
	id := c.Param("id")
	name := c.MustGet("userName").(string)
	if err := interfaceService.DeleteInterface(id, name); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(true, c)
}
