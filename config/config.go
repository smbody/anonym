package config

import (
	"github.com/spf13/viper"
)

//Init initializing new configuration file
func Init() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetDefault("host", "")
	viper.SetDefault("port", "8808")
	_ = viper.ReadInConfig()
}

func DataSourceName() string {
	return GetString("database.uri")
}

func DatabaseName () string {
	return GetString("database.name")
}

//GetString return string parameter value from config by his name
func GetString(name string) string {
	return viper.GetString(name)
}

//GetInteger return integer parameter value from config by his name
func GetInteger(name string) int {
	return viper.GetInt(name)
}
