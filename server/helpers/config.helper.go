package helpers

import (
	"github.com/spf13/viper"
)

func LoadENV() (err error) {
	viper.New()
	viper.SetConfigFile("../.env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	return viper.ReadInConfig()
}
