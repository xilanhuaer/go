package main

import (
	"interface/global"
	"interface/router"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"

)

func main() {
	global.Connection("./config/config.yaml")
	c := cors.Default()
	r := gin.Default()
	r.Use(c)
	router.Register(r)
	if global.DB != nil {
		db, _ := global.DB.DB()
		defer db.Close()
	}
	r.Run("0.0.0.0:80")
}
