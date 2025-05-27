package api

import "swiftDaily_myself/service"

type ApiGroup struct {
	UserApi
	BaseApi
}

var ApiGroupApp = new(ApiGroup)

var baseService = service.ServiceGroupApp.BaseService
