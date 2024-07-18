package initialize

import (
	"fmt"
	"gproject/internal/initialize/global"

	"github.com/spf13/viper"
)

func InitConFig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("fail to read viper config %v", err)
		return
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Config error %v", err)
		return
	}
}