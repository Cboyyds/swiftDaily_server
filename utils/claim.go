package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net"
	"swiftDaily_myself/global"
	"swiftDaily_myself/model/request"
)

func setCookie(c *gin.Context, tokenName, token string, maxAge int, host string) {
	if net.ParseIP(host) != nil {
		c.SetCookie(tokenName, token, maxAge, "/", "", false, true)
	} else {
		c.SetCookie(tokenName, token, maxAge, "/", host, false, true)
	}
}
func SetRefreshToken(c *gin.Context, token string, maxAge int) {
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	setCookie(c, "x-refresh-token", token, maxAge, host)
}
func ClearRefreshToken(c *gin.Context) {
	host, _, err := net.SplitHostPort(c.Request.Host) // 从请求头中的host中提取host和port
	if err != nil {
		host = c.Request.Host
	}
	setCookie(c, "x-refresh-token", "", -1, host)
}

func GetAccessToken(c *gin.Context) string {
	token := c.Request.Header.Get("x-access-token")
	return token
}

func GetRefreshToken(c *gin.Context) string {
	token, _ := c.Cookie("x-refresh-token")
	return token
}

func GetAccessClaims(c *gin.Context) (*request.JwtCustomClaims, error) {
	token := GetAccessToken(c)
	j := NewJWT()
	claims, err := j.ParshAccessToken(token)
	if err != nil {
		global.Log.Error("access token error", zap.Error(err))
	}
	return claims, err
}

func GetRefreshClaims(c *gin.Context) (*request.JwtCustomRefreshClaims, error) {
	token := GetRefreshToken(c)
	j := NewJWT()
	claims, err := j.ParshRefreshToken(token)
	if err != nil {
		global.Log.Error("refresh token error", zap.Error(err))
	}
	return claims, err
}

// 解析出用户信息claims
func GetUserInfo(c *gin.Context) *request.JwtCustomClaims {
	if claims, exists := c.Get("claims"); !exists { // 从上下文中获取claims
		if cl, err := GetAccessClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*request.JwtCustomClaims)
		return waitUse
	}
}

// 解析出用户的ID
func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetAccessClaims(c); err != nil {
			return 0
		} else {
			return cl.UserID
		}
	} else {
		return claims.(*request.JwtCustomClaims).UserID
	}
}

func GetUUID(c *gin.Context) uuid.UUID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetAccessClaims(c); err != nil {
			return uuid.Nil
		} else {
			return cl.UUID
		}
	} else {
		return claims.(*request.JwtCustomClaims).UUID
	}
}

func GetRoleID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetAccessClaims(c); err != nil {
			return 0
		} else {
			return cl.RoleID
		}
	} else {
		return claims.(*request.JwtCustomClaims).RoleID
	}
}
