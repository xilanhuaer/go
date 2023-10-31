package middleware

import (
	"interface/global"
	"interface/model/common/response"
	"strings"

	"github.com/gin-gonic/gin"
)

type JWTAuthMiddleware struct {
}

func (j *JWTAuthMiddleware) JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		urlPath := c.Request.URL.Path
		switch urlPath {
		case "/v1/user/register", "/v1/user/login":
			return
		}
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.FailWithMessage("请求头中auth为空", c)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.FailWithMessage("请求头中auth格式有误", c)
			c.Abort()
			return
		}
		mc, err := global.ParseJwt(parts[1])
		if err != nil {
			response.FailWithMessage("无效的Token", c)
			c.Abort()
			return
		}
		c.Set("userId", mc.UserId)
		c.Set("username", mc.UserName)
		c.Next()
	}
}
