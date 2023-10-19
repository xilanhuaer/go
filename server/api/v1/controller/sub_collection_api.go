package controller

import (
	"interface/model/common/response"
	"interface/model/entity"
	"log"

	"github.com/gin-gonic/gin"
)

type SubCollectionApi struct {
}

func (s *SubCollectionApi) CreateSubCollection(c *gin.Context) {
	var e entity.SubCollection
	err := c.ShouldBindJSON(&e)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = subCollectionService.CreateSubCollection(e)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(true, c)
}

func (s *SubCollectionApi) FindSubCollections(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	page_size := c.DefaultQuery("page_size", "10")
	params := map[string]string{
		"name":               c.DefaultQuery("name", ""),
		"enabled":            c.DefaultQuery("enabled", "1"),
		"created_at":         c.DefaultQuery("created_at", ""),
		"main_collection_id": c.DefaultQuery("main_collection_id", ""),
	}
	e, count, err := subCollectionService.FindSubCollections(page, page_size, params)
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
func (s *SubCollectionApi) FindSubCollection(c *gin.Context) {
	id := c.Param("id")
	data, err := subCollectionService.FindSubCollection(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(data, c)
}
func (s *SubCollectionApi) UpdateSubCollection(c *gin.Context) {
	id := c.Param("id")
	var e entity.SubCollection
	if err := c.ShouldBindJSON(&e); err != nil {
		log.Println(1)
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := subCollectionService.UpdateSubCollection(id, e)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(true, c)
}
func (s *SubCollectionApi) CheckSubCollectionEnable(c *gin.Context) {
	id := c.Param("id")
	var e entity.SubCollection
	if err := c.ShouldBindJSON(&e); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := subCollectionService.CheckSubCollectionEnable(id, e)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(true, c)
}
func (s *SubCollectionApi) DeleteSubCollection(c *gin.Context) {
	id := c.Param("id")
	if err := subCollectionService.DeleteSubCollection(id); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(true, c)
}
