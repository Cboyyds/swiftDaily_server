package core

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"os"
	"swiftDaily_myself/config"
	"swiftDaily_myself/global"
)

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

func InitConfig() *config.Config {
	filepath := "config.yaml"
	viper.SetConfigFile(filepath)
	// 获取文件MD5
	configMD5, err := ReadFileMD5(filepath)
	if err != nil {
		global.Log.Error("ReadFileMD5 error", zap.Error(err))
	}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config File:", err)
	}
	config := config.Config{}
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
