package controller

import (
	"interface/model/common/response"
	"interface/model/entity"
	"log"

	"github.com/gin-gonic/gin"
)

type SubCollectionController struct {
}

func (s *SubCollectionController) Create(context *gin.Context) {
	var e entity.SubCollection
	err := context.ShouldBindJSON(&e)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	err = subCollectionService.CreateSubCollection(e)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(true, context)
}

func (s *SubCollectionController) List(context *gin.Context) {
	page := context.DefaultQuery("page", "1")
	page_size := context.DefaultQuery("page_size", "10")
	params := map[string]string{
		"name":               context.DefaultQuery("name", ""),
		"enabled":            context.DefaultQuery("enabled", "1"),
		"created_at":         context.DefaultQuery("created_at", ""),
		"main_collection_id": context.DefaultQuery("main_collection_id", ""),
	}
	e, count, err := subCollectionService.FindSubCollections(page, page_size, params)
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
func (s *SubCollectionController) Find(context *gin.Context) {
	id := context.Param("id")
	data, err := subCollectionService.FindSubCollection(id)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(data, context)
}
func (s *SubCollectionController) Update(context *gin.Context) {
	id := context.Param("id")
	var e entity.SubCollection
	if err := context.ShouldBindJSON(&e); err != nil {
		log.Println(1)
		response.FailWithMessage(err.Error(), context)
		return
	}
	err := subCollectionService.UpdateSubCollection(id, e)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(true, context)
}
func (s *SubCollectionController) Enable(context *gin.Context) {
	id := context.Param("id")
	var e entity.SubCollection
	if err := context.ShouldBindJSON(&e); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	err := subCollectionService.CheckSubCollectionEnable(id, e)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(true, context)
}
func (s *SubCollectionController) Delete(context *gin.Context) {
	id := context.Param("id")
	if err := subCollectionService.DeleteSubCollection(id); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(true, context)
}
