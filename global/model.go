package global

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
	// DeletedAt 使用 gorm.DeletedAt 类型实现软删除功能，包含以下特性:
	// 1. 类型为 gorm.DeletedAt，本质上是 sql.NullTime 的别名
	// 2. 记录删除时间而非直接删除数据
	// 3. 自动填充删除时间戳
	// 4. 查询时默认过滤已删除记录(需使用 Unscoped() 查询已删除记录)
	// 5. `gorm:"index"` 为该字段创建索引提高查询效率
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
