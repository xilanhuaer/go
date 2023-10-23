package main

import (
	"interface/global"
	"interface/router"

	"github.com/gin-gonic/gin"
)

func main() {
	global.Connection("./config/config.yaml")
	r := gin.Default()
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))
	router.Register(r)
	if global.DB != nil {
		db, _ := global.DB.DB()
		defer db.Close()
	}
	r.Run("0.0.0.0:80")
}
