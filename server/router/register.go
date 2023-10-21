package router

import (
	"interface/api"

	"github.com/gin-gonic/gin"
)

func Register(route *gin.Engine) {
	jwtApi := api.ApiGroupApp.ControllerApiGroup.JWTAuthMiddlewareApi
	mainCollectionGroup := route.Group("/v1/collection/main")
	mainCollectionApi := api.ApiGroupApp.ControllerApiGroup.MainCollectionApi
	{
		mainCollectionGroup.POST("/", mainCollectionApi.CreateMainCollection)
		mainCollectionGroup.GET("/", mainCollectionApi.FindMainCollections)
		mainCollectionGroup.GET("/:id", mainCollectionApi.FindMainCollection)
		mainCollectionGroup.PUT("/:id", mainCollectionApi.UpdateMainCollection)
		mainCollectionGroup.PUT("/enable/:id", mainCollectionApi.CheckMainCollectionEnable)
		mainCollectionGroup.DELETE("/:id", mainCollectionApi.DeleteMainCollection)
	}
	subCollectionGroup := route.Group("/v1/collection/sub")
	subCollectionApi := api.ApiGroupApp.ControllerApiGroup.SubCollectionApi

	{
		subCollectionGroup.POST("/", subCollectionApi.CreateSubCollection)
		subCollectionGroup.GET("/", subCollectionApi.FindSubCollections)
		subCollectionGroup.GET("/:id", subCollectionApi.FindSubCollection)
		subCollectionGroup.PUT("/:id", subCollectionApi.UpdateSubCollection)
		subCollectionGroup.PUT("/enable/:id", subCollectionApi.CheckSubCollectionEnable)
		subCollectionGroup.DELETE("/:id", subCollectionApi.DeleteSubCollection)
	}
	interfaceGroup := route.Group("/v1/interface")
	interfaceApi := api.ApiGroupApp.ControllerApiGroup.InterfaceApi
	{
		interfaceGroup.POST("/", interfaceApi.CreateInterface)
		interfaceGroup.GET("/", jwtApi.JWTAuthMiddleware(), interfaceApi.FindInterfaces)
		interfaceGroup.GET("/:id", interfaceApi.FindInterface)
		interfaceGroup.PUT("/:id", interfaceApi.UpdateInterface)
		interfaceGroup.PUT("/enable/:id", interfaceApi.CheckInterfaceEnable)
		interfaceGroup.DELETE("/:id", interfaceApi.DeleteInterface)
	}
	interfaceImplGroup := route.Group("/v1/interface/impl")
	interfaceImplApi := api.ApiGroupApp.ControllerApiGroup.InterfaceImplApi
	{
		interfaceImplGroup.POST("/", interfaceImplApi.CreateImpl)
	}
}
