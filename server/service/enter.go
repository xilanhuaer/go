package service

import (
	"interface/service/assert"
	"interface/service/impl"
	"interface/service/json"
)

type ServiceGroup struct {
	ImplServiceGroup     impl.ServiceGroup
	AssertServiceGroup   assert.AssertServiceGroup
	JsonPathServiceGroup json.JsonPathServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
