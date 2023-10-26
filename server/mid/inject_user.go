package mid

import (
	"interface/global"

	"github.com/gin-gonic/gin"
)

// 增加用户信息注入
func InjectUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取用户信息
		username, exists := c.Get("username")
		if exists {
			global.DB.Set("Creator", username)
			global.DB.Set("Updator", username)
		}
		c.Next()
	}
}