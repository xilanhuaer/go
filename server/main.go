package main

import (
	"interface/global"
	"interface/router"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	global.Connection("./config/config.yaml")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://111.230.89.67"},
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Register(r)
	if global.DB != nil {
		db, _ := global.DB.DB()
		defer db.Close()
	}
	r.Run(":80")
}
