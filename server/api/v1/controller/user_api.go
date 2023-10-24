package controller

import (
	"encoding/json"
	"interface/model/common/response"
	"interface/model/entity"
	"log"

	"github.com/gin-gonic/gin"
)

// Path: go/server/api/v1/controller/user.go
type UserApi struct {
}

// 用户注册
func (ua *UserApi) UserRegister(c *gin.Context) {
	var (
		u    entity.User
		data map[string]interface{}
	)
	b, _ := c.GetRawData()
	_ = json.Unmarshal(b, &data)
	if value, ok := data["code"]; ok {
		if value != "register_code" {
			response.FailWithMessage("验证码错误", c)
			return
		}
	} else {
		response.FailWithMessage("请输入验证码", c)
		return
	}

	// 获取用户信息
	if err := c.ShouldBindJSON(&u); err != nil {
		response.FailWithMessage(err.Error(), c)
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

// 查看用户信息
func (ua *UserApi) UserInfo(c *gin.Context) {
	id := c.MustGet("userId").(uint)
	log.Println(id)
	userinfo, err := userService.FindUserInfo(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(userinfo, c)
}

// 修改密码
func (ua *UserApi) UpdatePassword(c *gin.Context) {
	var (
		editPassword entity.EditPassword
		userId       = c.MustGet("userId").(uint)
	)
	if err := c.ShouldBindJSON(&editPassword); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userService.UpdatePassword(userId, editPassword); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithMessage("修改成功", c)
}
