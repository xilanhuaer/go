package service

import (
	"interface/service/assert"
	"interface/service/impl"
)

type ServiceGroup struct {
	ImplServiceGroup   impl.ServiceGroup
	AssertServiceGroup assert.AssertServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
