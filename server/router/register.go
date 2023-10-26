package router

import (
	"interface/api"
	"interface/middleware"

	"github.com/gin-gonic/gin"
)

func Register(route *gin.Engine) {
	jwt := &middleware.JWTAuthMiddleware{}
	userGroup := route.Group("/v1/user")
	userApi := api.ApiGroupApp.ControllerApiGroup.UserApi
	{
		userGroup.POST("/register", userApi.UserRegister)
		userGroup.POST("/login", userApi.UserLogin)
		userGroup.GET("/info", jwt.JWTAuthMiddleware(), userApi.UserInfo)
		userGroup.POST("/edit_password", jwt.JWTAuthMiddleware(), userApi.UpdatePassword)
	}
	mainCollectionGroup := route.Group("/v1/collection/main", jwt.JWTAuthMiddleware())
	mainCollectionApi := api.ApiGroupApp.ControllerApiGroup.MainCollectionApi
	{
		mainCollectionGroup.POST("", mainCollectionApi.CreateMainCollection)
		mainCollectionGroup.GET("/", mainCollectionApi.FindMainCollections)
		mainCollectionGroup.GET("/:id", mainCollectionApi.FindMainCollection)
		mainCollectionGroup.PUT("/:id", mainCollectionApi.UpdateMainCollection)
		mainCollectionGroup.PUT("/enable/:id", mainCollectionApi.CheckMainCollectionEnable)
		mainCollectionGroup.DELETE("/:id", mainCollectionApi.DeleteMainCollection)
	}
	subCollectionGroup := route.Group("/v1/collection/sub", jwt.JWTAuthMiddleware())
	subCollectionApi := api.ApiGroupApp.ControllerApiGroup.SubCollectionApi

	{
		subCollectionGroup.POST("", subCollectionApi.CreateSubCollection)
		subCollectionGroup.GET("/", subCollectionApi.FindSubCollections)
		subCollectionGroup.GET("/:id", subCollectionApi.FindSubCollection)
		subCollectionGroup.PUT("/:id", subCollectionApi.UpdateSubCollection)
		subCollectionGroup.PUT("/enable/:id", subCollectionApi.CheckSubCollectionEnable)
		subCollectionGroup.DELETE("/:id", subCollectionApi.DeleteSubCollection)
	}
	interfaceGroup := route.Group("/v1/interface", jwt.JWTAuthMiddleware())
	interfaceApi := api.ApiGroupApp.ControllerApiGroup.InterfaceApi
	{
		interfaceGroup.POST("", interfaceApi.CreateInterface)
		interfaceGroup.GET("/", interfaceApi.FindInterfaces)
		interfaceGroup.GET("/:id", interfaceApi.FindInterface)
		interfaceGroup.PUT("/:id", interfaceApi.UpdateInterface)
		interfaceGroup.PUT("/enable/:id", interfaceApi.CheckInterfaceEnable)
		interfaceGroup.DELETE("/:id", interfaceApi.DeleteInterface)
	}
	interfaceImplGroup := route.Group("/v1/interface/impl", jwt.JWTAuthMiddleware())
	interfaceImplApi := api.ApiGroupApp.ControllerApiGroup.InterfaceImplApi
	{
		interfaceImplGroup.POST("", interfaceImplApi.CreateImpl)
		interfaceImplGroup.GET("/", interfaceImplApi.FindInterfaceImplements)
		interfaceImplGroup.GET("/:id", interfaceImplApi.FindInterfaceImplById)
		interfaceImplGroup.PUT("/:id", interfaceImplApi.UpdateInterfaceImplById)
	}
	requestController := api.ApiGroupApp.ControllerApiGroup.RequestController
	requestGroup := route.Group("/v1/request", jwt.JWTAuthMiddleware())
	{
		requestGroup.GET("/", requestController.TestRequest)
	}
}
