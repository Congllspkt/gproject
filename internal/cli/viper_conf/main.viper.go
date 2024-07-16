package viper_conf

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Databases []struct {
		User string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host string `mapstructure:"host"`
		DbName string `mapstructure:"dbName"`
	} `mapstructure:"databases"`
}

func GetViperConf() {

	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("fail to read viper config %v", err)
		return
	}

	fmt.Println("port: ", viper.GetInt("server.port"))
	fmt.Println("jwt: ", viper.GetString("security.jwt.key"))

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Config error %v", err)
		return
	}

	fmt.Println("config port: ", config.Server.Port)
	for _, db := range config.Databases {
		fmt.Printf("databases User: %s, password: %s, host: %s, dbName: %s.", db.User, db.Password, db.Host, db.DbName)
		fmt.Println()
	}
















}