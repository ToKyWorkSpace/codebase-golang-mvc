package config

import "github.com/spf13/viper"

func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}

func GetEnv(keys string) string {
	return viper.GetString(keys)
}
