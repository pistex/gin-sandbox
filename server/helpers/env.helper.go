package helpers

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

func ENVGetString(name string) string {
	v := viper.GetString(name)
	if v == "" {
		return os.Getenv(name)
	}

	return v
}

func ENVGetInt(name string) int {
	v := viper.GetInt(name)
	if v == 0 {
		s := os.Getenv(name)
		v, err := strconv.Atoi(s)
		if err != nil {
			return 0
		}
		return v
	}

	return v
}
