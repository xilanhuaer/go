package controller

import (
	"interface/utils/request"

	"github.com/gin-gonic/gin"
)

type RequestController struct {
}

func (rc *RequestController) TestRequest(c *gin.Context) {
	request.NewRequest("/v1/interface")
}
