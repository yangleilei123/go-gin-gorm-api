package config

import (
	"log"

	"github.com/spf13/viper" // 配置读取库
)

var Conf *viper.Viper // 全局配置变量

// InitConfig 加载配置文件 config.yaml
func InitConfig() {
	Conf = viper.New()
	Conf.SetConfigName("config") // 文件名（不含扩展名）
	Conf.SetConfigType("yaml")   // 配置格式
	Conf.AddConfigPath(".")      // 配置路径

	// 读取配置文件失败会直接退出程序
	if err := Conf.ReadInConfig(); err != nil {
		log.Fatalf("Config read error: %v", err)
	}
}
