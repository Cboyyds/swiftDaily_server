package database

import (
	"github.com/google/uuid"
	"time"
)

// Login登录日志表
type Login struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	UserID      uint      `json:"user_id"`
	UUID        uuid.UUID `json:"uuid" gorm:"type:char(36)"`
	LoginMethod string    `json:"login_method" gorm:"type:varchar(50)"`
	IP          string    `json:"ip" gorm:"type:varchar(50)"`
	Address     string    `json:"address" gorm:"type:varchar(255)"`
	OS          string    `json:"os" gorm:"type:varchar(50)"`
	DeviceInfo  string    `json:"device_info" gorm:"type:varchar(255)"`
	BrowserInfo string    `json:"browser_info" gorm:"type:varchar(255)"`
	Status      int       `json:"status" gorm:"enum(1,2,3);default:1"` // 1,在线，2，离线，3，冻结
	CreateAt    time.Time `json:"create_at" gorm:"type:datetime;default:1;CURRENT_TIMESTAMP;comment:'登录时间'"`
	OutAt       time.Time `json:"out_at" gorm:"type:datetime;comment:'登出时间;default:NULL'"`
}
