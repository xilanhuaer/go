package router

import (
	"interface/api"

	"github.com/gin-gonic/gin"
)

func Register(route *gin.Engine) {
	userGroup := route.Group("/v1/user")
	userApi := api.ApiGroupApp.ControllerApiGroup.UserController
	{
		userGroup.POST("/register", userApi.Register)
		userGroup.POST("/login", userApi.Login)
		userGroup.GET("/info", userApi.UserInfo)
		userGroup.POST("/edit_password", userApi.UpdatePassword)
	}
	mainCollectionGroup := route.Group("/v1/collection/main")
	mainCollectionController := api.ApiGroupApp.ControllerApiGroup.MainCollectionController
	{
		mainCollectionGroup.POST("", mainCollectionController.Create)
		mainCollectionGroup.GET("/", mainCollectionController.List)
		mainCollectionGroup.GET("/:id", mainCollectionController.Find)
		mainCollectionGroup.PUT("/:id", mainCollectionController.Update)
		mainCollectionGroup.PUT("/enable/:id", mainCollectionController.Enable)
		mainCollectionGroup.DELETE("/:id", mainCollectionController.Delete)
	}
	subCollectionGroup := route.Group("/v1/collection/sub")
	subCollectionApi := api.ApiGroupApp.ControllerApiGroup.SubCollectionController

	{
		subCollectionGroup.POST("", subCollectionApi.Create)
		subCollectionGroup.GET("/", subCollectionApi.List)
		subCollectionGroup.GET("/:id", subCollectionApi.Find)
		subCollectionGroup.PUT("/:id", subCollectionApi.Update)
		subCollectionGroup.PUT("/enable/:id", subCollectionApi.Enable)
		subCollectionGroup.DELETE("/:id", subCollectionApi.Delete)
	}
	interfaceGroup := route.Group("/v1/interface")
	interfaceController := api.ApiGroupApp.ControllerApiGroup.InterfaceController
	{
		interfaceGroup.POST("", interfaceController.Create)
		interfaceGroup.GET("/", interfaceController.List)
		interfaceGroup.GET("/:id", interfaceController.Find)
		interfaceGroup.PUT("/:id", interfaceController.Update)
		interfaceGroup.PUT("/enable/:id", interfaceController.Enable)
		interfaceGroup.DELETE("/:id", interfaceController.Delete)
	}
	interfaceImplGroup := route.Group("/v1/interface/impl")
	interfaceImplController := api.ApiGroupApp.ControllerApiGroup.InterfaceImplController
	{
		interfaceImplGroup.POST("", interfaceImplController.Create)
		interfaceImplGroup.GET("/", interfaceImplController.List)
		interfaceImplGroup.GET("/:id", interfaceImplController.Find)
		interfaceImplGroup.PUT("/:id", interfaceImplController.Update)
	}
	requestController := api.ApiGroupApp.ControllerApiGroup.RequestController
	requestGroup := route.Group("/v1/request")
	{
		requestGroup.GET("/test", requestController.TestRequest)
	}
}
