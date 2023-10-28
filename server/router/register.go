package router

import (
	"interface/api"
	"interface/middleware"

	"github.com/gin-gonic/gin"
)

func Register(route *gin.Engine) {
	jwt := &middleware.JWTAuthMiddleware{}
	userGroup := route.Group("/v1/user")
	userApi := api.ApiGroupApp.ControllerApiGroup.UserController
	{
		userGroup.POST("/register", userApi.Register)
		userGroup.POST("/login", userApi.Login)
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
	interfaceController := api.ApiGroupApp.ControllerApiGroup.InterfaceController
	{
		interfaceGroup.POST("", interfaceController.Create)
		interfaceGroup.GET("/", interfaceController.List)
		interfaceGroup.GET("/:id", interfaceController.Find)
		interfaceGroup.PUT("/:id", interfaceController.Update)
		interfaceGroup.PUT("/enable/:id", interfaceController.Enable)
		interfaceGroup.DELETE("/:id", interfaceController.Delete)
	}
	interfaceImplGroup := route.Group("/v1/interface/impl", jwt.JWTAuthMiddleware())
	interfaceImplController := api.ApiGroupApp.ControllerApiGroup.InterfaceImplController
	{
		interfaceImplGroup.POST("", interfaceImplController.CreateImpl)
		interfaceImplGroup.GET("/", interfaceImplController.FindInterfaceImplements)
		interfaceImplGroup.GET("/:id", interfaceImplController.FindInterfaceImplById)
		interfaceImplGroup.PUT("/:id", interfaceImplController.UpdateInterfaceImplById)
	}
	requestController := api.ApiGroupApp.ControllerApiGroup.RequestController
	requestGroup := route.Group("/v1/request", jwt.JWTAuthMiddleware())
	{
		requestGroup.GET("/", requestController.TestRequest)
	}
}
