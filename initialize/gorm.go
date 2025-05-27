package initialize

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"swiftDaily_myself/global"
)

func InitGorm() *gorm.DB {
	mysqlCfg := global.Config.Mysql
	db, err := gorm.Open(mysql.Open(mysqlCfg.DSN()), &gorm.Config{
		// 设置日志级别
		Logger: logger.Default.LogMode(mysqlCfg.LogLevel()),
	})
	if err != nil {
		global.Log.Error("Failed to connect to Mysql", zap.Error(err))
		os.Exit(1)
	}
	// // 获取底层的sql数据库链接对象，用于配置链接池
	sqlDB, err := db.DB()
	if err != nil {
		global.Log.Error("Failed to get sqlDB")
		os.Exit(1)
	}
	sqlDB.SetMaxIdleConns(mysqlCfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(mysqlCfg.MaxOpenConns)
	return db
}
