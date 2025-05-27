package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"swiftDaily_myself/config"
)

var (
	Config *config.Config
	Log    *zap.Logger
	DB     *gorm.DB
)
