package impl

import (
	"fmt"
	"interface/global"
	"interface/model/entity"
	"interface/utils"
)

type UserService struct {
}

// Path: go/server/service/impl/user_service.go
// 用户注册
func (u *UserService) Register(user *entity.User) (err error) {
	// 获取用户信息
	var (
		isValid bool
	)
	// 使用账号查询用户是否存在
	if err = global.DB.Where("account=?", user.Account).First(&user).Error; err == nil {
		return err
	}
	// 验证账号、密码、邮箱、电话号码格式
	isValid = utils.CheckAccount(user.Account)
	if !isValid {
		return fmt.Errorf("账号格式错误")
	}
	isValid = utils.CheckPassword(user.Password)
	if !isValid {
		return fmt.Errorf("密码格式错误")
	}
	isValid = utils.CheckEmail(user.Email)
	if !isValid {
		return fmt.Errorf("邮箱格式错误")
	}
	isValid = utils.CheckPhone(user.Phone)
	if !isValid {
		return fmt.Errorf("电话号码格式错误")
	}
	// 对密码进行加密
	user.Password = utils.SHA256V(user.Password)
	// 创建用户
	if err = global.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// 修改密码
// func (u *UserService) ChangePassword(user *entity.User, c *gin.Context) (err error) {
// 	// 获取用户信息
// 	var (
// 		isValid bool
// 	)
// 	// 使用账号查询用户是否存在
// 	if err = global.DB.Where("account=?", user.Account).First(&user).Error; err != nil {
// 		return err
// 	}
// 	// 验证密码格式
// 	isValid = utils.CheckPassword(user.Password)
// 	if !isValid {
// 		return fmt.Errorf("密码格式错误")
// 	}
// 	// 创建用户
// 	if err = global.DB.Create(&user).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// 用户登录
func (us *UserService) Login(user entity.User) (userinfo entity.UserInfo, err error) {
	var (
		u entity.User
	)
	if err = global.DB.Where("account=?", user.Account).First(&u).Error; err != nil {
		return userinfo, err
	}
	if u.Password != utils.SHA256V(user.Password) {
		return userinfo, fmt.Errorf("密码错误")
	}
	token, err := global.GenJwt(u.Id, u.Account)
	if err != nil {
		return userinfo, err
	}
	userinfo = entity.UserInfo{
		Id:      u.Id,
		Account: u.Account,
		Name:    u.Name,
		Email:   u.Email,
		Phone:   u.Phone,
		Token:   token,
	}
	return userinfo, nil
}
