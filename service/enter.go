package service

type ServiceGroup struct {
	JwtService
	BaseService
}

var ServiceGroupApp = new(ServiceGroup)
