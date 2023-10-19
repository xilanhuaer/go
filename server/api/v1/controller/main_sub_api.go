package controller

import (
	"interface/model/common/response"

	"github.com/gin-gonic/gin"
)

type MainSubApi struct {
}

func (msa *MainSubApi) MainSubList(c *gin.Context) {
	data, err := mainSubService.MainSubList()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.OKWithData(data, c)
}
