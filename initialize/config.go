package initialize

import (
	"fmt"
	"gin-xutao/global"
	"gin-xutao/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Configs()  {
	v := viper.New()
	v.SetConfigFile(utils.ConfigFile)
	v.SetConfigType("yaml")

	err := v.ReadInConfig()


	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err != nil {
		fmt.Println(err)
	}


	err = v.Unmarshal(&global.GVA_CONFIG)
	if err != nil {
		fmt.Println(err)
	}

}