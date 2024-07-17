package viper_conf

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
}