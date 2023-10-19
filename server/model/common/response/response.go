package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

const (
	SUCCESS = 200
	ERROR   = 0
)

func Result(code int, data interface{}, message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Data:    data,
		Message: message,
	})
}

func OK(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "success", c)
}
func OKWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}
func OKWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "success", c)
}
func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "error", c)
}
func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}
func FailWithDetail(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
