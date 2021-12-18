package redis

import (
	"github.com/go-redis/redis/v8"
	"os"
)

func Setup() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.ExpandEnv("${REDIS_HOST}:6379"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
