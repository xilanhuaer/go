package controller

import (
	"interface/service"
)

type ApiGroup struct {
	MainCollectionController
	SubCollectionController
	InterfaceController
	MainSubApi
	InterfaceImplController
	UserController
	RequestController
}

var (
	mainCollectionService = service.ServiceGroupApp.ImplServiceGroup.MainCollectionService
	subCollectionService  = service.ServiceGroupApp.ImplServiceGroup.SubCollectionService
	interfaceService      = service.ServiceGroupApp.ImplServiceGroup.InterfaceService
	mainSubService        = service.ServiceGroupApp.ImplServiceGroup.MainSubService
	interfaceImplService  = service.ServiceGroupApp.ImplServiceGroup.InterfaceImplService
	userService           = service.ServiceGroupApp.ImplServiceGroup.UserService
	requestService        = service.ServiceGroupApp.ImplServiceGroup.RequestService
)
