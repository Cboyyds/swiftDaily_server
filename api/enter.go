package api

import "swiftDaily_myself/service"

type ApiGroup struct {
	UserApi
	BaseApi
}

var ApiGroupApp = new(ApiGroup)

var baseService = service.ServiceGroupApp.BaseService // 虽然是小写，但是这个是api包，是引入了service的变量
// var userService = service.ServiceGroupApp.UserService
var gaodeService = service.ServiceGroupApp.GaodeService
var jwtService = service.ServiceGroupApp.JwtService
