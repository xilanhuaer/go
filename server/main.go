package main

import (
	"interface/global"
	"interface/middlewares"
	"interface/router"

	"github.com/gin-gonic/gin"
)

func main() {
	global.Connection("./config/config.yaml")
	r := gin.Default()
	r.Use(middlewares.Cors())
	router.Register(r)
	if global.DB != nil {
		db, _ := global.DB.DB()
		defer db.Close()
	}
	r.Run("0.0.0.0:80")
}
