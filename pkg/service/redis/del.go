package redis_client

import "context"

func (r redisClient) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
