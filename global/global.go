package global

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"swiftDaily_myself/config"
)

var (
	Config *config.Config
	Log    *zap.Logger
	DB     *gorm.DB
	Redis  *redis.Client
)
