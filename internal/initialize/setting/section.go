package setting

type Config struct {
	MySQL  MySQLSetting  `mapstructure:"mySQL"`
	Logger LoggerSetting `mapstructure:"logger"`
}

type MySQLSetting struct {
	Host           string `mapstructure:"host"`
	Port           int    `mapstructure:"port"`
	Username       string `mapstructure:"username"`
	Password       string `mapstructure:"password"`
	Dbname         string `mapstructure:"dbname"`
	MaxIdleConns   int    `mapstructure:"maxIdleConns"`
	MaxOpenConns   int    `mapstructure:"maxOpenConns"`
	ConMaxLifeTime int    `mapstructure:"conMaxLifeTime"`
}

type LoggerSetting struct {
	Log_level     string `mapstructure:"log_level"`
	File_log_name string `mapstructure:"file_log_name"`
	Max_size      int    `mapstructure:"max_size"`
	Max_backups   int    `mapstructure:"max_backups"`
	Max_age       int    `mapstructure:"max_age"`
	Compress      bool   `mapstructure:"compress"`
}