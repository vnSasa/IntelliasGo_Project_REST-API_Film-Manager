package app

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func GetRedisConn() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.Addr"),
		Password: "",
		DB:       0,
	})
}
