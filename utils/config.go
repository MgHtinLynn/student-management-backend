package utils

import "github.com/spf13/viper"

func LoadConfig(key string) string {
	viper.SetConfigFile(".env")

	_ = viper.ReadInConfig()

	return viper.Get(key).(string)
}
