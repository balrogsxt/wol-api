package app

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// InitConfig 初始化配置文件
func InitConfig() error {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	if err := v.Unmarshal(&Config); err != nil {
		return err
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		if err := v.Unmarshal(&Config); err != nil {
			Logger.Warn("配置文件重载失败: " + err.Error())
		} else {
			Logger.Info("配置文件已重载")
		}
	})

	return nil
}

type AppConfig struct {
	Debug  bool   //是否开启调试模式
	JwtKey string `mapstructure:"jwt_key"`
	ApiKey string `mapstructure:"api_key"`
	Http   struct {
		Host string
		Port int
	}
	Wol struct {
		MacAddress string `mapstructure:"mac_addr"` //目标机器mac地址
		Ip         string //目标机器Ip地址
		Network    string //目标机器网卡名称
		Port       int    //休眠端口
	}
}
