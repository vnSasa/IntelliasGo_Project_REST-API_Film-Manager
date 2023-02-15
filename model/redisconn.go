package app

import (
	"github.com/go-redis/redis/v8"
)

func GetRedisConn() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
