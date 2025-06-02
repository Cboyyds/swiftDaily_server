package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"swiftDaily_myself/global"
	"swiftDaily_myself/model/request"
	"time"
)

type JWT struct {
	AccessTokenSecret  []byte
	RefreshTokenSecret []byte
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotVaildYet = errors.New("token not vaild yet")
	TokenMalformed   = errors.New("token malformed")
	TokenInvaild     = errors.New("could't handle this token")
)

func NewJWT() *JWT {
	return &JWT{
		AccessTokenSecret:  []byte(global.Config.Jwt.AccessTokenSecret),
		RefreshTokenSecret: []byte(global.Config.Jwt.RefreshTokenSecret),
	}
}

func (j *JWT) CreateAccessClaims(baseClaims request.BaseClaims) request.JwtCustomClaims {
	ep, _ := ParseDuration(global.Config.Jwt.AccessTokenExpire)
	claim := request.JwtCustomClaims{
		BaseClaims: baseClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"TAP"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),
			Issuer:    global.Config.Jwt.Issuer,
		},
	}
	return claim
}

func (j *JWT) CreateAccessToken(claim request.JwtCustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(j.AccessTokenSecret)
}

func (j *JWT) CreateRefreshClaims(baseClaims request.BaseClaims) request.JwtCustomRefreshClaims {
	ep, _ := ParseDuration(global.Config.Jwt.RefreshTokenExpire) // 获取过期时间
	claim := request.JwtCustomRefreshClaims{
		UserID: baseClaims.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"TAP"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),
			Issuer:    global.Config.Jwt.Issuer,
		},
	}
	return claim
}
func (j *JWT) CreateRefreshToken(claim request.JwtCustomRefreshClaims) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return refreshToken.SignedString(j.RefreshTokenSecret)
}
func (j *JWT) ParshAccessToken(tokenString string) (*request.JwtCustomClaims, error) {
	claims, err := j.parseToken(tokenString, &request.JwtCustomClaims{}, j.AccessTokenSecret)
	if err != nil {
		return nil, err
	}
	if customClaims, ok := claims.(*request.JwtCustomClaims); ok {
		return customClaims, nil
	}
	return nil, TokenInvaild
}
func (j *JWT) ParshRefreshToken(tokenString string) (*request.JwtCustomRefreshClaims, error) {
	refreshClaims, err := j.parseToken(tokenString, &request.JwtCustomRefreshClaims{}, j.RefreshTokenSecret)
	if err != nil {
		return nil, err
	}
	if refreshClaims, ok := refreshClaims.(*request.JwtCustomRefreshClaims); ok {
		return refreshClaims, nil
	}
	return nil, TokenInvaild
}

// 解析access和refresh的通用方法
func (j *JWT) parseToken(tokenString string, claims jwt.Claims, secretKey interface{}) (interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			switch {
			case ve.Errors&jwt.ValidationErrorMalformed != 0:
				return nil, TokenMalformed
			case ve.Errors&jwt.ValidationErrorExpired != 0:
				return nil, TokenExpired
			case ve.Errors&jwt.ValidationErrorNotValidYet != 0:
				return nil, TokenNotVaildYet
			default:
				return nil, TokenInvaild
			}
		}
		return nil, TokenInvaild
	}
	if token.Valid {
		return token.Claims, nil
	}
	return nil, TokenInvaild
}
