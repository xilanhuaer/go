package main

import (
	"interface/global"
	"interface/middleware"
	"interface/router"

	"github.com/gin-gonic/gin"
)

func main() {
	global.GetConfig("./config/config.yaml")
	global.Connection()
	r := gin.New()
	r.Use(gin.Logger())
	jwt := &middleware.JWTAuthMiddleware{}
	r.Use(jwt.JWTAuthMiddleware())
	router.Register(r)
	if global.DB != nil {
		db, _ := global.DB.DB()
		defer db.Close()
	}
	r.Run(":80")
}
