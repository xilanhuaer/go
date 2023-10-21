package utils

import "regexp"

// path: go/server/utils/format_check.go

// 验证账号
// 账号格式：字母开头，允许5-16字节，允许字母数字下划线
func CheckAccount(account string) bool {
	pattern := "^[a-zA-Z][a-zA-Z0-9_]{4,15}$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(account)
}

// 验证密码
// 密码格式：以字母开头，长度在6~18之间，只能包含字母、数字和下划线
func CheckPassword(password string) bool {
	pattern := "^[a-zA-Z]\\w{5,17}$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(password)
}

// 验证邮箱
// 邮箱格式：合法邮箱
func CheckEmail(email string) bool {
	pattern := "^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// 验证电话号码]
// 电话号码格式：11位手机号码
func CheckPhone(phone string) bool {
	pattern := "^1[3456789]\\d{9}$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(phone)
}
