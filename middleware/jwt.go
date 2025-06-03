package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"swiftDaily_myself/global"
	"swiftDaily_myself/model/database"
	"swiftDaily_myself/model/request"
	"swiftDaily_myself/model/response"
	"swiftDaily_myself/service"
	"swiftDaily_myself/utils"
)

var jwtService = service.ServiceGroupApp.JwtService

// 配置jwt中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := utils.GetAccessToken(c)
		refreshToken := utils.GetRefreshToken(c)
		// 黑名单
		if jwtService.IsInBlacklist(refreshToken) {
			utils.ClearRefreshToken(c)
			response.NoAuth("The user does not exist", c)
			c.Abort()
			return
		}
		//
		j := utils.NewJWT()
		claims, err := j.ParshAccessToken(accessToken)
		if err != nil { // error
			if accessToken == "" || errors.Is(err, utils.TokenExpired) { //  accessToken过期或者为空
				refreshClaims, err := j.ParshRefreshToken(refreshToken)
				if err != nil {
					utils.ClearRefreshToken(c)
					response.FailWithMessage("refreshToken is expired", c)
					c.Abort()
					return
				}
				var user database.User
				if err := global.DB.Select("uuid", "role_id").Find(&user, refreshClaims.UserID).Error; err != nil {
					utils.ClearRefreshToken(c)
					response.NoAuth("The user does not exist", c)
					c.Abort()
					return
				}
				newAccessClaims := j.CreateAccessClaims(request.BaseClaims{
					UserID: refreshClaims.UserID,
					UUID:   user.UUID,
					RoleID: user.RoleID,
				})
				newAccessToken, err := j.CreateAccessToken(newAccessClaims)
				if err != nil {
					utils.ClearRefreshToken(c)
					response.NoAuth("The user does not exist", c)
					c.Abort()
					return
				}
				c.Header("new-access-token", newAccessToken)
				c.Header("new-access-expire-at", strconv.FormatInt(newAccessClaims.ExpiresAt.Unix(), 10))
				c.Set("claims", newAccessClaims)
				c.Next()
				return
			}
			utils.ClearRefreshToken(c)
			response.NoAuth("Invaild accessToken", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
