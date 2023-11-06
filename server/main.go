package main

import (
	"interface/global"
	"interface/middleware"
	"interface/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	global.GetConfig("./config/config.yaml")
	global.Connection()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	jwt := &middleware.JWTAuthMiddleware{}
	r.Use(jwt.JWTAuthMiddleware())
	router.Register(r)
	if global.DB != nil {
		db, _ := global.DB.DB()
		defer db.Close()
	}
	r.Run(":80")
}
