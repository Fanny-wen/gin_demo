package viper

import (
	"fmt"
	"github.com/spf13/viper"
)

// 常量
const (
	AppConfigFile   = "./Assets/configs/app.yaml"   // 配置文件
	MysqlConfigFile = "./Assets/configs/mysql.yaml" // 配置文件
)

var (
	MysqlConfig  map[string]interface{}
	AppConfig    map[string]interface{}
	ServerConfig map[string]interface{}
)


// 应用信息
type app struct {
	Desc       string `yaml:"desc"`
	Addr       string `yaml:"addr"`
	ConfigFile string `yaml:"configFile"`
	Version    string `yaml:"version"`
	Env        string `yaml:"env"`
}

// MySQL信息
type mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func init() {
	AppConfig = InitConfig(AppConfigFile)
	MysqlConfig = InitConfig(MysqlConfigFile)
	fmt.Println(AppConfig)
	fmt.Println(MysqlConfig)
	ServerConfig = map[string]interface{}{
		"App": AppConfig["app"],
		"Mysql": MysqlConfig["mysql"],
	}
}

// InitConfig 初始化viper配置解析包，函数可接受命令行参数
func InitConfig(f string) (config map[string]interface{}) {
	var configFile = f
	if len(configFile) == 0 {
		// 读取默认配置文件
		panic("配置文件不能为空！")
	}
	// 读取配置文件
	v := viper.New()
	v.SetConfigFile(f)      // 指定配置文件路径
	v.SetConfigType("yaml") // 如果配置文件的名称中没有扩展名，则需要配置此项

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("配置解析失败:%s\n", err))
	}
	// 动态监测配置文件
	//v.WatchConfig()
	//v.OnConfigChange(func(in fsnotify.Event) {
	//	fmt.Println("配置文件发生改变")
	//	if err := v.Unmarshal(&global.GvaConfig); err != nil {
	//		panic(fmt.Errorf("配置重载失败:%s\n", err))
	//	}
	//})
	//if err := v.Unmarshal(&global.GvaConfig); err != nil {
	//	panic(fmt.Errorf("配置重载失败:%s\n", err))
	//}
	//// 设置配置文件
	//global.GvaConfig.App.ConfigFile = configFile
	return v.AllSettings()
}
