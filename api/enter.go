package api

type ApiGroup struct {
	UserApi
	BaseApi
}

var ApiGroupApp = new(ApiGroup)
