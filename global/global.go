package global

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"swiftDaily_myself/config"
)

var (
	Config     *config.Config
	Log        *zap.Logger
	DB         *gorm.DB
	Redis      *redis.Client
	Ctx        = context.Background()
	BlackCache local_cache.Cache
)
