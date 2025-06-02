package initialize

import (
	"github.com/redis/go-redis/v9"
	"swiftDaily_myself/global"
)

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Addr,
		Password: global.Config.Redis.Password,
		DB:       global.Config.Redis.DB,
	})
	return rdb
}
