package impl

import (
	"fmt"
	"interface/global"
	"interface/model/entity"
	"interface/utils"
)

type UserService struct {
}

// 用户注册
func (u *UserService) Register(user *entity.User) (err error) {
	var (
		userinfo entity.User
		isValid  bool
	)
	if err = global.DB.Where("account=?", user.Account).First(&userinfo).Error; err != nil {
		isValid = utils.CheckAccount(user.Account)
		if !isValid {
			return fmt.Errorf("账号格式错误")
		}
		isValid = utils.CheckPassword(user.Password)
		if !isValid {
			return fmt.Errorf("密码格式错误")
		}
		user.Password = utils.SHA256V(user.Password)
		isValid = utils.CheckEmail(user.Email)
		if !isValid {
			return fmt.Errorf("邮箱格式错误")
		}
		isValid = utils.CheckPhone(user.Phone)
		if !isValid {
			return fmt.Errorf("手机号格式错误")
		}
		err = global.DB.Create(&user).Error
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("用户已存在")
}

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
