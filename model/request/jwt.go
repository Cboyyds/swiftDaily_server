package request

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type BaseClaims struct {
	UserID uint
	RoleID uint
	UUID   uuid.UUID
}
type JwtCustomClaims struct {
	BaseClaims
	jwt.RegisteredClaims
}
type JwtCustomRefreshClaims struct {
	UserID uint
	jwt.RegisteredClaims
}
