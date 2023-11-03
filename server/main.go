package main

import (
	"interface/global"
	"interface/middleware"
	"interface/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	global.Connection("./config/config.yaml")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.Default())
	jwt := &middleware.JWTAuthMiddleware{}
	r.Use(jwt.JWTAuthMiddleware())
	router.Register(r)
	if global.DB != nil {
		db, _ := global.DB.DB()
		defer db.Close()
	}
	r.Run(":80")
}
