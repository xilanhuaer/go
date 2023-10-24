package main

import (
	"interface/global"
	"interface/router"

	"github.com/gin-gonic/gin"
)

func main() {
	global.Connection("./config/config.yaml")
	r := gin.New()
	r.Use(gin.Logger())
	router.Register(r)
	if global.DB != nil {
		db, _ := global.DB.DB()
		defer db.Close()
	}
	r.Run(":80")
}
