package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
)

func main() {
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

	viper.SetConfigName("config") //  设置配置文件名 (不带后缀)
	viper.AddConfigPath("/etc/appname/")   // 第一个搜索路径
	viper.AddConfigPath("$HOME/.appname")  // 可以多次调用添加路径
	viper.AddConfigPath(".")               // 比如添加当前目录
	err := viper.ReadInConfig() // 搜索路径，并读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}




