package helpers

import (
	"github.com/spf13/viper"
)

func LoadENV(path string) (err error) {
	viper.New()
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	return viper.ReadInConfig()
}
