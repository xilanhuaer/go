package api

import "interface/api/v1/controller"

type ApiGroup struct {
	ControllerApiGroup controller.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
