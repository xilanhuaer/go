package controller

import (
	"interface/model/common/response"
	"interface/model/entity"
	"log"

	"github.com/gin-gonic/gin"
)

type MainCollectionApi struct{}

func (m *MainCollectionApi) CreateMainCollection(c *gin.Context) {
	var mainCollection entity.MainCollection
	err := c.ShouldBindJSON(&mainCollection)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = mainCollectionService.CreateMainCollection(mainCollection)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(true, c)
}

func (e *MainCollectionApi) FindMainCollections(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	page_size := c.DefaultQuery("page_size", "10")
	params := map[string]string{
		"name":       c.DefaultQuery("name", ""),
		"enabled":    c.DefaultQuery("enabled", "1"),
		"created_at": c.DefaultQuery("created_at", ""),
	}
	maincollections, count, err := mainCollectionService.FindMainCollections(page, page_size, params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data := response.PageResult{
		List:  maincollections,
		Total: count,
	}
	response.OKWithData(data, c)
}
func (e *MainCollectionApi) FindMainCollection(c *gin.Context) {
	id := c.Param("id")
	data, err := mainCollectionService.FindMainCollection(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(data, c)
}
func (e *MainCollectionApi) UpdateMainCollection(c *gin.Context) {
	id := c.Param("id")
	var mainCollection entity.MainCollection
	if err := c.ShouldBindJSON(&mainCollection); err != nil {
		log.Println(1)
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := mainCollectionService.UpdateMainCollection(id, mainCollection)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(true, c)
}
func (e *MainCollectionApi) CheckMainCollectionEnable(c *gin.Context) {
	id := c.Param("id")
	var mainCollection entity.MainCollection
	if err := c.ShouldBindJSON(&mainCollection); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := mainCollectionService.CheckMainCollectionEnable(id, mainCollection)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(true, c)
}
func (e *MainCollectionApi) DeleteMainCollection(c *gin.Context) {
	id := c.Param("id")
	if err := mainCollectionService.DeleteMainCollection(id); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(true, c)
}
