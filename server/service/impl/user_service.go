package impl

import (
	"fmt"
	"interface/global"
	"interface/model/entity"
	"interface/utils"

	"github.com/gin-gonic/gin"
)

type UserService struct {
}

// Path: go/server/service/impl/user_service.go
// 用户注册
func (u *UserService) Register(user *entity.User, c *gin.Context) (err error) {
	// 获取用户信息
	var (
		userInfo entity.User
		isValid  bool
	)
	if err = c.ShouldBindJSON(&userInfo); err != nil {
		return err
	}
	// 使用账号查询用户是否存在
	if err = global.DB.Where("account=?", userInfo.Account).First(&user).Error; err != nil {
		return err
	}
	// 验证账号、密码、邮箱、电话号码格式
	isValid = utils.CheckAccount(userInfo.Account)
	if !isValid {
		return fmt.Errorf("账号格式错误")
	}
	isValid = utils.CheckPassword(userInfo.Password)
	if !isValid {
		return fmt.Errorf("密码格式错误")
	}
	isValid = utils.CheckEmail(userInfo.Email)
	if !isValid {
		return fmt.Errorf("邮箱格式错误")
	}
	isValid = utils.CheckPhone(userInfo.Phone)
	if !isValid {
		return fmt.Errorf("电话号码格式错误")
	}
	// 创建用户
	if err = global.DB.Create(&userInfo).Error; err != nil {
		return err
	}
	return nil
}

// 修改密码
func (u *UserService) ChangePassword(user *entity.User, c *gin.Context) (err error) {
	// 获取用户信息
	var (
		userInfo entity.User
		isValid  bool
	)
	if err = c.ShouldBindJSON(&userInfo); err != nil {
		return err
	}
	// 使用账号查询用户是否存在
	if err = global.DB.Where("account=?", userInfo.Account).First(&user).Error; err != nil {
		return err
	}
	// 验证密码格式
	isValid = utils.CheckPassword(userInfo.Password)
	if !isValid {
		return fmt.Errorf("密码格式错误")
	}
	// 创建用户
	if err = global.DB.Create(&userInfo).Error; err != nil {
		return err
	}
	return nil
}
