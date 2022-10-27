package utils

import "github.com/spf13/viper"

func LoadConfig(key string) string {
	viper.AutomaticEnv()

	_ = viper.ReadInConfig()

	return viper.Get(key).(string)
}
