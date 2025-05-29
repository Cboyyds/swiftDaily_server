package service

type ServiceGroup struct {
	JwtService
	BaseService
	GaodeService
}

var ServiceGroupApp = new(ServiceGroup)
