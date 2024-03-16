package config

import (
	"fmt"
	"fun-with-gin/pkg/utils"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"os"
	"time"
)

const (
	REDIS_HOST      = "REDIS_HOST"
	REDIS_PORT      = "REDIS_PORT"
	REDIS_PASSWORD  = "REDIS_PASSWORD"
	REDIS_DB        = "REDIS_DB"
	REDIS_TIMEOUT   = "REDIS_TIMEOUT"
	MAX_REQUEST     = "MAX_REQUEST"
	EXPIRED_REQUEST = "EXPIRED_REQUEST"
)

type RedisConfig struct {
	Host           string
	Port           string
	DB             int
	Password       string
	Timeout        time.Duration
	MaxRequest     int
	ExpiredRequest time.Duration
}

func Redis() RedisConfig {
	return RedisConfig{
		Host:           os.Getenv(REDIS_HOST),
		Port:           os.Getenv(REDIS_PORT),
		DB:             utils.ConvertInt(os.Getenv(REDIS_DB)),
		Password:       os.Getenv(REDIS_PASSWORD),
		Timeout:        utils.DefaultValueDuration(REDIS_TIMEOUT, "10s"),
		MaxRequest:     utils.ConvertInt(os.Getenv(MAX_REQUEST)),
		ExpiredRequest: utils.DefaultValueDuration(EXPIRED_REQUEST, "10s"),
	}
}

func NewRedisConnection() *redis.Client {
	config := Redis()
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})
	rdb.WithTimeout(config.Timeout)

	_, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		panic(errors.Wrap(err, "redis"))
	}

	return rdb
}
