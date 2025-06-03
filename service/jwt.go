package service

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"swiftDaily_myself/global"
	"swiftDaily_myself/model/database"
	"swiftDaily_myself/utils"
)

type JwtService struct {
}

// SetRedisJwt设置根据refreshtoken是否过期设置
func (jwtService *JwtService) SetRedisJWT(jwt string, uuid uuid.UUID, ctx context.Context) error {
	dr, err := utils.ParseDuration(global.Config.Jwt.RefreshTokenExpire)
	if err != nil {
		return err
	}
	return global.Redis.Set(ctx, uuid.String(), jwt, dr).Err()
}

func (jwtService *JwtService) GetRedisJWT(uuid uuid.UUID, ctx context.Context) (string, error) {
	return global.Redis.Get(ctx, uuid.String()).Result()
}

func (jwtService *JwtService) JoinInBlackList(jwtList database.JwtBlacklist) error {
	// 将jwt记录插入到数据库中的黑名单
	if err := global.DB.Create(&jwtList).Error; err != nil {
		return err
	}
	// 将jwt添加到内存中的黑名单缓存
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return nil
}

func (jwtService *JwtService) IsInBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
}

func LoadAll() {
	var data []string
	if err := global.DB.Model(&database.JwtBlacklist{}).Pluck("jwt", &data).Error; err != nil {
		global.Log.Error("Failed to load Jwt blacklist from the database", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	}
}
