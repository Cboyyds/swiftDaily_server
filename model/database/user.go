package database

import (
	"github.com/google/uuid"
	"swiftDaily_myself/global"
)

type User struct {
	global.Model
	UUID      uuid.UUID `json:"uuid" gorm:"type:char(36);unique"`
	Password  string    `json:"-"` // gorm:"type:json"`,这样的写法也会被认为是json字段也是能进行反序列化，这里对字符的判断是交给前端做，密码传到后端判断是否正确
	UserName  string    `json:"user_name"`
	Signature string    `json:"signature"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"` // 头像链接
	// JobNumber   string    `json:"job_number"`
	CompanyID uint   `json:"company_ID"`
	DeptID    uint   `json:"dept_id"`
	RoleID    uint   `json:"role_id"`
	Role      Role   `json:"user" gorm:"foreignKey:ID"`
	Status    int    `json:"status" gorm:"enum(1,2,3);default:2"` // 1,在线，2，离线，3，冻结
	DeletedBy string `json:"deleted_by"`                          // 是谁删除了他的人的id
}
