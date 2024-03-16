package redis_client

import (
	"fun-with-gin/pkg/domain/service"
	"github.com/go-redis/redis/v8"
)

type redisClient struct {
	client *redis.Client
}

func NewRedisClient(client *redis.Client) service.RedisClient {
	return &redisClient{
		client: client,
	}
}
