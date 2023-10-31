package controller

import (
	"interface/model/common/response"
	"strings"

	"github.com/gin-gonic/gin"
)

type RequestController struct {
}

func (rc *RequestController) TestRequest(c *gin.Context) {
	body := "{\"account\":\"xilanhua\",\"password\":\"xilanhua655\"}"
	code, resBody, err := requestService.Request("POST", "http://111.230.89.67:82/v1/user/login", nil, nil, strings.NewReader(body))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	var res struct {
		code int
		data string
	}
	res.code = code
	res.data = resBody
	response.OKWithData(res, c)
}
