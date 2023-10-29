package controller

import (
	"interface/model/common/response"
	"interface/model/entity"
	"log"

	"github.com/gin-gonic/gin"
)

type InterfaceController struct {
}

func (interfaceController *InterfaceController) Create(c *gin.Context) {
	var (
		e        entity.Interface
		username = c.MustGet("username").(string)
	)
	err := c.ShouldBindJSON(&e)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	e.Creator = username
	e.Updator = username
	err = interfaceService.Create(e)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(true, c)
}

func (interfaceController *InterfaceController) List(c *gin.Context) {
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
	e, count, err := interfaceService.List(page, page_size, params)
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
func (interfaceController *InterfaceController) Find(c *gin.Context) {
	id := c.Param("id")
	data, err := interfaceService.Find(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(data, c)
}
func (i *InterfaceController) Update(c *gin.Context) {
	id := c.Param("id")
	name := c.MustGet("username").(string)
	var e entity.Interface
	e.Updator = name
	if err := c.ShouldBindJSON(&e); err != nil {
		log.Println(1)
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := interfaceService.Update(id, e)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(true, c)
}
func (interfaceController *InterfaceController) Enable(c *gin.Context) {
	id := c.Param("id")
	name := c.MustGet("username").(string)
	var e entity.Interface
	if err := c.ShouldBindJSON(&e); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := interfaceService.Enable(id, e.Enabled, name)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(true, c)
}
func (interfaceController *InterfaceController) Delete(c *gin.Context) {
	id := c.Param("id")
	name := c.MustGet("username").(string)
	if err := interfaceService.Delete(id, name); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(true, c)
}
