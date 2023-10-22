package controller

import (
	"bytes"
	"interface/global"
	"interface/model/entity"

	"github.com/gin-gonic/gin"
)

type RequestController struct {
}

func (rc *RequestController) TestRequest(c *gin.Context) {
	jsonData := `{"account":"xilanhua","password":"xilanhua666"}"}`
	postRequest := entity.NewRequest("POST", global.HOST+"/api/v1/user/login")
	postRequest.SetJsonBody(jsonData)
	postRequest.SetHeader("Content-Type", "application/json")
	resp, err := postRequest.Send()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}
	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  buf.String(),
	})
}