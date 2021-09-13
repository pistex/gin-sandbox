package helpers

import (
	"github.com/spf13/viper"
)

func LoadENV(filename string) (err error) {
	viper.New()
	viper.SetConfigFile(filename)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	return viper.ReadInConfig()
}
