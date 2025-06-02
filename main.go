package main

import (
	"swiftDaily_myself/core"
	"swiftDaily_myself/global"
	"swiftDaily_myself/initialize"
)

func main() {
	global.Config = core.InitConfig() // config的配置里面是不能有golbal.Log的东西的，因为还未执行，因此会有空指针
	global.Log = core.InitLogger()
	global.DB = initialize.InitGorm()
	global.Redis = initialize.InitRedis()
	initialize.InitRouter()

	initialize.InitCorn()
	core.RunServer()
}
