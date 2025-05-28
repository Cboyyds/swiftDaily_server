v1.0
~搭建项目
    1,config.go配置初始化
    2,写入配置文件
    3，读取配置文件---用viper---用MD5解决编辑器会连续调用两次回调函数，输出两次文件改变

```
package config

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	System  System  `json:"system" yaml:"system"`
	Mysql   Mysql   `json:"mysql" yaml:"mysql"`
	Redis   Redis   `json:"redis" yaml:"redis"`
	Zap     Zap     `json:"zap" yaml:"zap"`
	Jwt     Jwt     `json:"jwt" yaml:"jwt"`
	Email   Email   `json:"email" yaml:"email"`
	Captcha Captcha `json:"captcha" yaml:"captcha"`
	Website Website `json:"website" yaml:"website"`
	Upload  Upload  `json:"upload" yaml:"upload"`
}

// 因为用viper进行监控文件变更，编译器的问题会连续调用2次回调函数，因此这里用md5解决问题
func GetMD5(s []byte) string {
	m := md5.New()
	m.Write(s)
	return hex.EncodeToString(m.Sum(nil))
}
func ReadFileMD5(sfile string) (string, error) {
	ssconfig, err := os.ReadFile(sfile)
	if err != nil {
		return "", err
	}
	return GetMD5(ssconfig), nil
}

// 用viper读取配置文件

func InitConfig() *Config {
	filepath := "config.yaml"
	viper.SetConfigFile(filepath)
	// 获取文件MD5
	configMD5, err := ReadFileMD5(filepath)
	if err != nil {
		log.Fatal(err)
	}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config File:", err)
	}
	var config Config = Config{}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("unable to decode into struct: %v", err)
	}
	
	// 在返回前继续监听配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 这里用md5解决编译器连续调用2次回调函数的问题
		tconfigMD5, err := ReadFileMD5(filepath)
		if err != nil {
			log.Fatal(err)
		}
		if tconfigMD5 == configMD5 {
			return
		}
		if err := viper.Unmarshal(&config); err != nil {
			log.Fatal(err)
		}
		configMD5 = tconfigMD5
	})
	
	// 返回后仍然保持监听，因为 WatchConfig 和 OnConfigChange 是异步操作，return会再一次执行
	return &config
}

```

4，日志配置 ，初始化----使用zap，lumberjack

用三方

```
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
```

还有corn的使用：是一个用于管理定时任务的调度器

```
package initialize

import (
	"github.com/robfig/cron/v3"
	"os"
	"server/global"
	"server/task"

	"go.uber.org/zap"
)

// ZapLogger 结构体实现了 cron.Logger 接口的 Info 和 Error 方法，这些方法用于接收 cron 包生成的日志并使用 zap 进行记录,将记录生成到文件里面，而不是控制台
type ZapLogger struct {
	logger *zap.Logger //
}

func (l *ZapLogger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Info(msg, zap.Any("keysAndValues", keysAndValues))
}

func (l *ZapLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	l.logger.Error(msg, zap.Error(err), zap.Any("keysAndValues", keysAndValues))
}

func NewZapLogger() *ZapLogger {
	return &ZapLogger{logger: global.Log}
}

// InitCron 初始化定时任务
func InitCron() {
	// 将 cron 包的日志记录转发到 zap 日志库中，实现统一的日志管理和记录
	// c := cron.New(cron.WithLogger(cron.DefaultLogger)) // 使用默认的log，输出和zap.logger日志输出位置不一样，因此需要自己重写一个
	c := cron.New(cron.WithLogger(NewZapLogger()))
	err := task.RegisterScheduledTasks(c)
	if err != nil {
		global.Log.Error("Error scheduling cron job:", zap.Error(err))
		os.Exit(1)
	}
	c.Start()
}


```

![image-20250526203547293](C:\Users\c博\AppData\Roaming\Typora\typora-user-images\image-20250526203547293.png)

5，gin框架的使用

```
go get github.com/gin-gonic/gin
```

6.使用请求和捕获错误中间件

```
package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http/httputil"
	"os"
	"strings"
	"swiftDaily_myself/global"
	"time"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		
		cost := time.Since(start)
		global.Log.Info(path,
			// 记录响应状态码
			zap.Int("status", c.Writer.Status()),
			// 请求方法
			zap.String("method", c.Request.Method),
			// 请求路径
			zap.String("path", path),
			// 请求参数
			zap.String("query", query),
			// 请求耗时
			zap.String("cost", cost.String()),
			// Ip
			zap.String("ip", c.ClientIP()),
			// user-agent信息
			zap.String("user-agent", c.Request.UserAgent()),
			// 错误信息
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
		)
	}
}

// GinRecovery 用于捕获和处理请求中的panic错误
// 该错误确保服务在遇到未处理的异常时不会崩溃，并通过日志系统提供详细的错误跟踪
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			// 检查是否发生了panic错误
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connect reset by peer") {
							brokenPipe = true
						}
					}
				}
				// 获取请求信息，包括请求体等
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				// 如果时broken pipe'，则只记录错误信息，不记录堆栈信息
				if brokenPipe {
					global.Log.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// 由于链接断开，不能再向客户端写入状态码
					_ = c.Error(err.(error))
					c.Abort()
					return
				}
				
				// 如果是其他类型的panic，根据stack参数决定是否记录堆栈信息
				if stack {
					// 记录详细的错误和堆栈信息
					global.Log.Error("[Recovery from panic ]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)))
				}
			}
			
		}()
		// 继续执行后续请求处理
		c.Next()
	}
}

```

7.配置gorm，使用gorm进行数据库的建表操作

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

```
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

```

声明Model：最好不直接用gorm.Model因为无法直接打tag，自己写也要用上gorm.DeleteAt,不然就没法软删除

```
type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
```

```
// DeletedAt 使用 gorm.DeletedAt 类型实现软删除功能，包含以下特性:
// 1. 类型为 gorm.DeletedAt，本质上是 sql.NullTime 的别名
// 2. 记录删除时间而非直接删除数据
// 3. 自动填充删除时间戳
// 4. 查询时默认过滤已删除记录(需使用 Unscoped() 查询已删除记录)
// 5. `gorm:"index"` 为该字段创建索引提高查询效率
```

创建用户表：

```
package database

import (
	"server/global"
	"server/model/appTypes"

	"github.com/gofrs/uuid"
)

// User 用户表
type User struct {
	global.MODEL
	UUID      uuid.UUID         `json:"uuid" gorm:"type:char(36);unique"`              // uuid
	Username  string            `json:"username"`                                      // 用户名
	Password  string            `json:"-"`                                             // "-" ，转化为json字段时会忽略字段，防止敏感信息在api中暴露                                          // 密码
	Email     string            `json:"email"`                                         // 邮箱
	Openid    string            `json:"openid"`                                        // openid
	Avatar    string            `json:"avatar" gorm:"size:255"`                        // 头像：邮箱注册的头像或 QQ 登录的空间头像
	Address   string            `json:"address"`                                       // 地址
	Signature string            `json:"signature" gorm:"default:'签名是空白的，这位用户似乎比较低调。'"` // 签名
	RoleID    appTypes.RoleID   `json:"role_id"`                                       // 角色 ID
	Register  appTypes.Register `json:"register"`                                      // 注册来源
	Freeze    bool              `json:"freeze"`                                        // 用户是否被冻结
}

```

推了一下代码库，然后创建一下自己的分支

8，路由配置





9，JWT配置

------

我突然感觉具体写到哪一步，做了什么，用了什么东西好像没有什么意义，就像一个说明书，

真正不会的地方可以通过提交历史进行观察，

从图片验证码获取开始

## 技术：

Cookic和session

```
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
```





我才发现原来邮箱验证码是我自己生成发给我自己的啊
