package controller

import (
	"interface/model/common/response"
	"interface/model/entity"
	"interface/orm"
	"log"

	"github.com/gin-gonic/gin"
)

type MainCollectionController struct{}

func (mainCollectionController *MainCollectionController) Create(context *gin.Context) {
	var (
		mainCollection entity.MainCollection
		username       = context.MustGet("username").(string)
	)
	err := context.ShouldBindJSON(&mainCollection)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	err = orm.Create(&mainCollection, username)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(true, context)
}

func (mainCollectionController *MainCollectionController) List(context *gin.Context) {
	page := context.DefaultQuery("page", "1")
	page_size := context.DefaultQuery("page_size", "10")
	params := map[string]string{
		"name":       context.DefaultQuery("name", ""),
		"enabled":    context.DefaultQuery("enabled", "1"),
		"created_at": context.DefaultQuery("created_at", ""),
	}
	maincollections, count, err := mainCollectionService.List(page, page_size, params)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	data := response.PageResult{
		List:  maincollections,
		Total: count,
	}
	response.OKWithData(data, context)
}
func (mainCollectionController *MainCollectionController) Find(context *gin.Context) {
	id := context.Param("id")
	data, err := mainCollectionService.Find(id)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(data, context)
}
func (mainCollectionController *MainCollectionController) Update(context *gin.Context) {
	var (
		mainCollection entity.MainCollection
		username       = context.MustGet("username").(string)
	)
	if err := context.ShouldBindJSON(&mainCollection); err != nil {
		log.Println(1)
		response.FailWithMessage(err.Error(), context)
		return
	}
	err := orm.Update(&mainCollection, username)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(true, context)
}
func (mainCollectionController *MainCollectionController) Enable(context *gin.Context) {

	var (
		mainCollection entity.MainCollection
		id             = context.Param("id")
		username       = context.MustGet("username").(string)
	)
	if err := context.ShouldBindJSON(&mainCollection); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	err := orm.Enable(&mainCollection, id, username)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(true, context)
}
func (mainCollectionController *MainCollectionController) Delete(context *gin.Context) {
	var (
		id       = context.Param("id")
		username = context.MustGet("username").(string)
	)
	if err := orm.Delete(&entity.MainCollection{}, id, username); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(true, context)
}
