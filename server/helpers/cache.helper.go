package helpers

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func NewCache() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", ENVGetString("CACHE_HOST"), ENVGetString("CACHE_PORT")),
		Password: ENVGetString("CACHE_PASSWORD"),
		DB:       ENVGetInt("CACHE_DB"),
	})

	return rdb
}
