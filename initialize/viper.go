package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"path/filepath"
	"time"
)

const GlobalConfigFile = "D:/GoProject/src/github.com/Fanny-wen/gin_demo/config.yaml" // 配置文件

// GC 全局配置
var GC = GlobalConfig{}

// GlobalConfig 配置信息
type GlobalConfig struct {
	App   App   `yaml:"app"`
	Mysql Mysql `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
}

// Redis
type Redis struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// 应用信息
type App struct {
	Desc       string `yaml:"desc"`
	Addr       string `yaml:"addr"`
	ConfigFile string `yaml:"configFile"`
	Version    string `yaml:"version"`
	Env        string `yaml:"env"`
}

// MySQL信息
type Mysql struct {
	Host      string `yaml:"host"`
	Port      int `yaml:"port"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	DbName    string `yaml:"dbname"`
	Charset   string `yaml:"charset"`
	ParseTime string `yaml:"parseTime"`
	Loc       string `yaml:"loc"`
}

// InitConfig 初始化viper配置解析包，函数可接受命令行参数
func InitConfig(f string, gc *GlobalConfig) {
	var configFile = f
	if len(configFile) == 0 {
		// 读取默认配置文件
		panic("配置文件不能为空！")
	}

	v := viper.New()

	// SetConfigName() +  AddConfigPath() = SetConfigFile() 所以一下两种方式二选一, 这里选择方式一
	// 方式一:
	v.SetConfigName(filepath.Base(f)) // 配置文件名
	v.AddConfigPath(filepath.Dir(f))  // 指定配置文件路径, 多次调用以添加多个搜索路径
	// 方式二:
	//v.SetConfigFile(f) // 指定配置文件路径

	v.SetConfigType("yaml") // 设置文件类型
	//查找并读取配置文件
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("配置解析失败:%s\n", err))
	}
	// viper 支持将配置Unmarshal到一个结构体中，为结构体中的对应字段赋值。
	_ = v.Unmarshal(gc)

	// 设置并覆盖配置值
	v.Set("TestPath", "1")

	go func() {
		for {
			time.Sleep(time.Duration(5) * time.Second)
			// 动态监测配置文件
			v.WatchConfig()
			v.OnConfigChange(func(in fsnotify.Event) {
				// viper配置发生变化了 执行响应的操作
				fmt.Println("配置文件发生改变")
				if err := v.Unmarshal(gc); err != nil {
					panic(fmt.Errorf("配置重载失败:%s\n", err))
				}
			})
		}
	}()
	// 设置配置文件
	gc.App.ConfigFile = configFile
}
