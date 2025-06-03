package database

import "swiftDaily_myself/global"

type JwtBlacklist struct {
	global.Model
	Jwt string `json:"jwt" gorm:"type:text"` // jwt
}
