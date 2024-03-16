package service

import (
	"context"
	"time"
)

type RedisClient interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value any, exp time.Duration) error
	Del(ctx context.Context, key string) error
}
