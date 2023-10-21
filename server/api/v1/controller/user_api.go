package controller

import (
	"interface/model/common/response"
	"interface/model/entity"

	"github.com/gin-gonic/gin"
)

// Path: go/server/api/v1/controller/user.go
type UserApi struct {
}

// 用户注册
func (ua *UserApi) UserRegister(c *gin.Context) {
	var (
		u entity.User
	)
	// 获取用户信息
	if err := c.ShouldBindJSON(&u); err == nil {
		response.FailWithMessage("账号已存在", c)
		return
	}
	// 注册
	if err := userService.Register(&u); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithMessage("注册成功", c)

}

// 用户登录
func (ua *UserApi) UserLogin(c *gin.Context) {
	var (
		u entity.User
	)
	// 获取用户信息
	if err := c.ShouldBindJSON(&u); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 登录
	userinfo, err := userService.Login(u)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(userinfo, c)
}
