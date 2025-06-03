package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"swiftDaily_myself/global"
	"swiftDaily_myself/model/database"
	"swiftDaily_myself/model/request"
	"swiftDaily_myself/model/response"
	"swiftDaily_myself/utils"
	"time"
)

type UserApi struct {
}

func (u *UserApi) EmailRegister(c *gin.Context) {
	var req request.Register
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Failed to bind JSON", c)
		global.Log.Error("Failed to bind JSON", zap.Error(err))
		return
	}
	// session := sessions.Default(c)
	// savedEmail := session.Get("email")
	// if savedEmail != nil || savedEmail.(string) == req.Email {
	// 	response.FailWithMessage("Email verification failed", c)
	// 	global.Log.Error("Email verification failed", zap.Error(errors.New("Email verification failed")))
	// 	return
	// }
	// savedCode := session.Get("verification_code")
	// if savedCode.(string) != req.VerificationCode {
	// 	response.FailWithMessage("verification code error", c)
	// 	global.Log.Error("verification code error", zap.Error(errors.New("verification code error")))
	// 	return
	// }
	// savedTime := session.Get("expire_time")
	// if savedTime.(int64) < time.Now().Unix() {
	// 	response.FailWithMessage("time expire", c)
	// 	global.Log.Error("time expire", zap.Error(errors.New("time expire")))
	// 	return
	// }
	var user database.User = database.User{
		Email:     req.Email,
		UserName:  req.Username,
		Password:  utils.BcryptHash(req.Password), // 对密码进行加密处理
		CompanyID: req.CompanyID,
		UUID:      uuid.Must(uuid.NewV6()),
	}
	if err := baseService.EmailRegister(user); err != nil {
		response.FailWithMessage(err.Error(), c)
		global.Log.Error(err.Error(), zap.Error(err))
		return
	}
	response.OK(c)
	global.Log.Info("register success", zap.Any("user", user))
	return
}
func (u *UserApi) EmailLogin(c *gin.Context) {
	var req request.Login
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Failed to bind JSON", c)
		global.Log.Error("Failed to bind JSON", zap.Error(err))
		return
	}
	//  判断验证码是否正确
	// if !store.Verify(req.CaptchaID, req.Captcha, true) {
	// 	response.FailWithMessage("验证码错误", c)
	// 	global.Log.Error("验证码错误", zap.Error(errors.New("验证码错误")))
	// 	return
	// }
	user := database.User{
		Email:    req.Account,
		Password: req.Password,
	}
	user, err := baseService.EmailLogin(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		global.Log.Error(err.Error(), zap.Error(err))
		return
	}
	// response.OKWithDetail(user, "登录成功", c)
	// global.Log.Info("login success", zap.Any("user", user))
	// 注册成功后，生成token并返回
	u.TokenNext(c, user)
}

func (userApi *UserApi) TokenNext(c *gin.Context, user database.User) {
	// 是否冻结
	// if user.Freeze {
	// 	response.FailWithMessage("用户被冻结", c)
	// 	return
	// }
	//
	baseClaims := request.BaseClaims{
		UserID: user.ID,
		UUID:   user.UUID,
		RoleID: user.RoleID,
	}
	j := utils.NewJWT()
	accessClaims := j.CreateAccessClaims(baseClaims)
	accessToken, err := j.CreateAccessToken(accessClaims)
	if err != nil {
		global.Log.Error("create access token error", zap.Error(err))
		response.FailWithMessage("create access token error", c)
		return
	}
	// 创建刷新令牌
	refreshClaim := j.CreateRefreshClaims(baseClaims)
	refreshToken, err := j.CreateRefreshToken(refreshClaim)
	if err != nil {
		global.Log.Error("create refresh token error", zap.Error(err))
		response.FailWithMessage("create refresh token error", c)
		return
	}
	// 检查用户jwt是否存在redis中
	jwtStr, err := jwtService.GetRedisJWT(user.UUID, global.Ctx)
	
	if errors.Is(err, redis.Nil) {
		userApi.SetRedisWithJwt(refreshToken, accessToken, refreshClaim, accessClaims, user, c)
	}
	if err == nil {
		// Redis 中已存在该用户的 JWT，将旧的 JWT 加入黑名单，并设置新的 token
		var blacklist database.JwtBlacklist
		blacklist.Jwt = jwtStr
		if err := jwtService.JoinInBlackList(blacklist); err != nil {
			global.Log.Error(err.Error(), zap.Error(err))
			response.FailWithMessage(err.Error(), c)
			return
		}
		userApi.SetRedisWithJwt(refreshToken, accessToken, refreshClaim, accessClaims, user, c)
	}
	if err != redis.Nil || err != nil {
		global.Log.Error("Failed to set login status", zap.Error(err))
		response.FailWithMessage("Failed to set login status", c)
	}
}

// 感觉应该写在service层
func (userApi *UserApi) SetRedisWithJwt(refreshToken string, accessToken string, refreshClaim request.JwtCustomRefreshClaims, accessClaims request.JwtCustomClaims, user database.User, c *gin.Context) {
	if err := jwtService.SetRedisJWT(refreshToken, user.UUID, global.Ctx); err != nil {
		global.Log.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 设置刷新令牌并返回
	utils.SetRefreshToken(c, refreshToken, int(refreshClaim.ExpiresAt.Unix()-time.Now().Unix()))
	c.Set("user_id", user.ID)
	response.OKWithDetail(response.Login{
		User:              user,
		AccessToken:       accessToken,
		AccessTokenExpire: accessClaims.ExpiresAt.Unix() * 1000, // 该代码将 accessClaims.ExpiresAt 的 Unix 时间戳（秒级）转换为毫秒级，并赋值给 AccessTokenExpire 字段。
	}, "Successful login", c)
}
