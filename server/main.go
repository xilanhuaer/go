package main

import (
	"interface/global"
	"interface/router"

	"github.com/gin-gonic/gin"
)

func main() {
	global.Connection("./config/config.yaml")
	// 跨域
	Cors()
	r := gin.Default()
	router.Register(r)
	if global.DB != nil {
		db, _ := global.DB.DB()
		defer db.Close()
	}
	r.Run("0.0.0.0:80")
}

func Cors() {
	panic("unimplemented")
}
