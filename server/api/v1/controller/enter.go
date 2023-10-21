package controller

import "interface/service"

type ApiGroup struct {
	MainCollectionApi
	SubCollectionApi
	InterfaceApi
	MainSubApi
	InterfaceImplApi
	JWTAuthMiddlewareApi
}

var (
	mainCollectionService = service.ServiceGroupApp.ImplServiceGroup.MainCollectionService
	subCollectionService  = service.ServiceGroupApp.ImplServiceGroup.SubCollectionService
	interfaceService      = service.ServiceGroupApp.ImplServiceGroup.InterfaceService
	mainSubService        = service.ServiceGroupApp.ImplServiceGroup.MainSubService
	interfaceImplService  = service.ServiceGroupApp.ImplServiceGroup.InterfaceImplService
)
