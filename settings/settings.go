package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name         string `mapstructure:"name"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	Mode         string `mapstructure:"mode"`
	*MysqlConfig `mapstructure:"mysql"`
	*LogConfig   `mapstructure:"logger"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}

type LogConfig struct {
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"maxsize"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Lever      string `mapstructure:"lever"`
}

var Conf = new(Config)

func Init() (err error) {
	viper.SetConfigFile("./config/config.yaml") // 制定配置文件路径
	err = viper.ReadInConfig()                  //读取配置文件
	if err != nil {
		fmt.Printf("ReadInConfig conf failed:%v\n", err)
		return err
	}
	// 反序列化成结构体
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("Unmarshal Config failed:%v\n", err)
		return err
	}
	//监听文件
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config Change")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("Unmarshal Config failed:%v\n", err)
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}
	})
	return
}
