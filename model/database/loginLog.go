package database

import (
	"swiftDaily_myself/global"
)

// Login登录日志表
type Login struct {
	global.Model
	UserID      uint   `json:"user_id"`
	User        User   `json:"user" gorm:"foreginKey:UseID"`
	LoginMethod string `json:"login_method"`
	IP          string `json:"ip"`
	Address     string `json:"address"`
	OS          string `json:"os"`
	DeviceInfo  string `json:"device_info"`
	BrowserInfo string `json:"browser_info"`
	Status      int    `json:"status"`
}
