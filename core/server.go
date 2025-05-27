package core

import (
	"go.uber.org/zap"
	"swiftDaily_myself/global"
	"swiftDaily_myself/initialize"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	addr := global.Config.System.Addr()
	s := initialize.InitRouter()
	//
	
	// 初始化服务器并启动
	if err := s.Run(addr); err != nil {
		global.Log.Error("server run failed", zap.Error(err))
	}
}
