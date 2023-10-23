package router

import (
	"interface/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Register(route *gin.Engine) {
	jwtApi := api.ApiGroupApp.ControllerApiGroup.JWTAuthMiddlewareApi
	userGroup := route.Group("/v1/user")
	userGroup.Use(cors.Default())
	userApi := api.ApiGroupApp.ControllerApiGroup.UserApi
	{
		userGroup.POST("/register", userApi.UserRegister)
		userGroup.POST("/login", userApi.UserLogin)
		userGroup.GET("/info", jwtApi.JWTAuthMiddleware(), userApi.UserInfo)
		userGroup.POST("/edit_password", jwtApi.JWTAuthMiddleware(), userApi.UpdatePassword)
	}
	mainCollectionGroup := route.Group("/v1/collection/main", jwtApi.JWTAuthMiddleware())
	mainCollectionGroup.Use(cors.Default())
	mainCollectionApi := api.ApiGroupApp.ControllerApiGroup.MainCollectionApi
	{
		mainCollectionGroup.POST("/", mainCollectionApi.CreateMainCollection)
		mainCollectionGroup.GET("/", mainCollectionApi.FindMainCollections)
		mainCollectionGroup.GET("/:id", mainCollectionApi.FindMainCollection)
		mainCollectionGroup.PUT("/:id", mainCollectionApi.UpdateMainCollection)
		mainCollectionGroup.PUT("/enable/:id", mainCollectionApi.CheckMainCollectionEnable)
		mainCollectionGroup.DELETE("/:id", mainCollectionApi.DeleteMainCollection)
	}
	subCollectionGroup := route.Group("/v1/collection/sub", jwtApi.JWTAuthMiddleware())
	subCollectionGroup.Use(cors.Default())
	subCollectionApi := api.ApiGroupApp.ControllerApiGroup.SubCollectionApi

	{
		subCollectionGroup.POST("/", subCollectionApi.CreateSubCollection)
		subCollectionGroup.GET("/", subCollectionApi.FindSubCollections)
		subCollectionGroup.GET("/:id", subCollectionApi.FindSubCollection)
		subCollectionGroup.PUT("/:id", subCollectionApi.UpdateSubCollection)
		subCollectionGroup.PUT("/enable/:id", subCollectionApi.CheckSubCollectionEnable)
		subCollectionGroup.DELETE("/:id", subCollectionApi.DeleteSubCollection)
	}
	interfaceGroup := route.Group("/v1/interface", jwtApi.JWTAuthMiddleware())
	interfaceGroup.Use(cors.Default())
	interfaceApi := api.ApiGroupApp.ControllerApiGroup.InterfaceApi
	{
		interfaceGroup.POST("/", interfaceApi.CreateInterface)
		interfaceGroup.GET("/", interfaceApi.FindInterfaces)
		interfaceGroup.GET("/:id", interfaceApi.FindInterface)
		interfaceGroup.PUT("/:id", interfaceApi.UpdateInterface)
		interfaceGroup.PUT("/enable/:id", interfaceApi.CheckInterfaceEnable)
		interfaceGroup.DELETE("/:id", interfaceApi.DeleteInterface)
	}
	interfaceImplGroup := route.Group("/v1/interface/impl", jwtApi.JWTAuthMiddleware())
	interfaceImplGroup.Use(cors.Default())
	interfaceImplApi := api.ApiGroupApp.ControllerApiGroup.InterfaceImplApi
	{
		interfaceImplGroup.POST("/", interfaceImplApi.CreateImpl)
	}
	requestController := api.ApiGroupApp.ControllerApiGroup.RequestController
	requestGroup := route.Group("/v1/request")
	requestGroup.Use(cors.Default())
	{
		requestGroup.GET("/", requestController.TestRequest)
	}
}
