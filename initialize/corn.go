package initialize

// 重写zap.Logger的方法
import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"swiftDaily_myself/global"
)

// ZapLogger 是一个结构体，用于包装 zap.Logger
type ZapLogger struct {
	logger *zap.Logger
}

// Info 方法用于记录信息级别的日志
func (l *ZapLogger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Info(msg, zap.Any("keysAndValues", keysAndValues))
}

// Error 方法用于记录错误级别的日志
func (l *ZapLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	l.logger.Error(msg, zap.Error(err), zap.Any("keysAndValues", keysAndValues))
}

// NewZapLogger 函数用于创建一个新的 ZapLogger 实例
func NewZapLogger() *ZapLogger {
	return &ZapLogger{logger: global.Log}
}

// InitCron 函数用于初始化定时任务
func InitCorn() {
	c := cron.New(cron.WithLogger(NewZapLogger()))
	// err := task.RegisterScheduledTasks(c)
	c.Start()
}
