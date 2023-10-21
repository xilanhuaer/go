package router

import (
	"interface/api"

	"github.com/gin-gonic/gin"
)

func Register(route *gin.Engine) {
	jwtApi := api.ApiGroupApp.ControllerApiGroup.JWTAuthMiddlewareApi
	userGroup := route.Group("/v1/user")
	userApi := api.ApiGroupApp.ControllerApiGroup.UserApi
	{
		userGroup.POST("/register", userApi.UserRegister)
		userGroup.POST("/login", userApi.UserLogin)
	}
	mainCollectionGroup := route.Group("/v1/collection/main")
	mainCollectionApi := api.ApiGroupApp.ControllerApiGroup.MainCollectionApi
	{
		mainCollectionGroup.POST("/", jwtApi.JWTAuthMiddleware(), mainCollectionApi.CreateMainCollection)
		mainCollectionGroup.GET("/", jwtApi.JWTAuthMiddleware(), mainCollectionApi.FindMainCollections)
		mainCollectionGroup.GET("/:id", jwtApi.JWTAuthMiddleware(), mainCollectionApi.FindMainCollection)
		mainCollectionGroup.PUT("/:id", jwtApi.JWTAuthMiddleware(), mainCollectionApi.UpdateMainCollection)
		mainCollectionGroup.PUT("/enable/:id", jwtApi.JWTAuthMiddleware(), mainCollectionApi.CheckMainCollectionEnable)
		mainCollectionGroup.DELETE("/:id", jwtApi.JWTAuthMiddleware(), mainCollectionApi.DeleteMainCollection)
	}
	subCollectionGroup := route.Group("/v1/collection/sub")
	subCollectionApi := api.ApiGroupApp.ControllerApiGroup.SubCollectionApi

	{
		subCollectionGroup.POST("/", jwtApi.JWTAuthMiddleware(), subCollectionApi.CreateSubCollection)
		subCollectionGroup.GET("/", jwtApi.JWTAuthMiddleware(), subCollectionApi.FindSubCollections)
		subCollectionGroup.GET("/:id", jwtApi.JWTAuthMiddleware(), subCollectionApi.FindSubCollection)
		subCollectionGroup.PUT("/:id", jwtApi.JWTAuthMiddleware(), subCollectionApi.UpdateSubCollection)
		subCollectionGroup.PUT("/enable/:id", jwtApi.JWTAuthMiddleware(), subCollectionApi.CheckSubCollectionEnable)
		subCollectionGroup.DELETE("/:id", jwtApi.JWTAuthMiddleware(), subCollectionApi.DeleteSubCollection)
	}
	interfaceGroup := route.Group("/v1/interface")
	interfaceApi := api.ApiGroupApp.ControllerApiGroup.InterfaceApi
	{
		interfaceGroup.POST("/", jwtApi.JWTAuthMiddleware(), interfaceApi.CreateInterface)
		interfaceGroup.GET("/", jwtApi.JWTAuthMiddleware(), interfaceApi.FindInterfaces)
		interfaceGroup.GET("/:id", jwtApi.JWTAuthMiddleware(), interfaceApi.FindInterface)
		interfaceGroup.PUT("/:id", jwtApi.JWTAuthMiddleware(), interfaceApi.UpdateInterface)
		interfaceGroup.PUT("/enable/:id", jwtApi.JWTAuthMiddleware(), interfaceApi.CheckInterfaceEnable)
		interfaceGroup.DELETE("/:id", jwtApi.JWTAuthMiddleware(), interfaceApi.DeleteInterface)
	}
	interfaceImplGroup := route.Group("/v1/interface/impl")
	interfaceImplApi := api.ApiGroupApp.ControllerApiGroup.InterfaceImplApi
	{
		interfaceImplGroup.POST("/", jwtApi.JWTAuthMiddleware(), interfaceImplApi.CreateImpl)
	}
}
