package service

import "interface/service/impl"

type ServiceGroup struct {
	ImplServiceGroup impl.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
