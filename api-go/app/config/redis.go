package config

import (
	"runtime"

	"github.com/gofiber/storage/redis/v3"
)

var Redis *redis.Storage

func InitRedis() {
	Redis = redis.New(redis.Config{
		Host:     GetEnvVariable("REDIS_URL"),
		Port:     6379,
		Password: GetEnvVariable("REDIS_PASSWORD"),
		Database: 0,
		PoolSize: 10 * runtime.GOMAXPROCS(0),
		Username: GetEnvVariable("REDIS_USER"),
	})
}

func GetRedis() *redis.Storage {
	return Redis
}
