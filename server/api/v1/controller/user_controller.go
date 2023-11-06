package controller

import (
	"interface/global"
	"interface/model/common/response"
	"interface/model/entity"
	"log"

	"github.com/gin-gonic/gin"
)

// Path: go/server/api/v1/controller/user.go
type UserController struct{}

// 用户注册
func (userCenter *UserController) Register(context *gin.Context) {
	var (
		user         entity.User
		registerData struct {
			Account  string
			Password string
			Code     string
			Name     string
			Email    string
			Phone    string
		}
	)
	if err := context.ShouldBindJSON(&registerData); err != nil {
		response.FailWithMessage("请求参数错误", context)
		return
	}
	if registerData.Code != global.Config.System.Code {
		response.FailWithMessage("邀请码错误", context)
		return
	}
	{
		user.Account = registerData.Account
		user.Password = registerData.Password
		user.Name = registerData.Name
		user.Email = registerData.Email
		user.Phone = registerData.Phone

	}
	// 注册
	if err := userService.Register(&user); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithMessage("注册成功", context)
}

// 用户登录
func (userController *UserController) Login(context *gin.Context) {
	var user entity.User
	// 获取用户信息
	if err := context.ShouldBindJSON(&user); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	// 登录
	userinfo, err := userService.Login(user)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(userinfo, context)
}

// 查看用户信息
func (userController *UserController) UserInfo(context *gin.Context) {
	id := context.MustGet("userId").(uint)
	log.Println(id)
	userinfo, err := userService.FindUserInfo(id)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(userinfo, context)
}

// 修改密码
func (userController *UserController) UpdatePassword(context *gin.Context) {
	var (
		editPassword entity.EditPassword
		userId       = context.MustGet("userId").(uint)
	)
	if err := context.ShouldBindJSON(&editPassword); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	if err := userService.UpdatePassword(userId, editPassword); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithMessage("修改成功", context)
}

// 个人中心
func (userController *UserController) UserCenter(context *gin.Context) {
	userId := context.MustGet("userId").(uint)
	userinfo, err := userService.FindUserInfo(userId)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(userinfo, context)
}

// 修改个人信息
func (userController *UserController) UpdateUserInfo(context *gin.Context) {
	var (
		userId = context.MustGet("userId").(uint)
		user   entity.User
	)
	if err := context.ShouldBindJSON(&user); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	if err := userService.UpdateUserInfo(userId, user); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithMessage("修改成功", context)
}
