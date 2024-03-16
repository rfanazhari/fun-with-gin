package redis_client

import (
	"context"
	"time"
)

func (r redisClient) Set(ctx context.Context, key string, value any, exp time.Duration) error {
	return r.client.Set(ctx, key, value, exp).Err()
}
