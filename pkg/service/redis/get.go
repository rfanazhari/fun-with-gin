package redis_client

import "context"

func (r redisClient) Get(ctx context.Context, key string) (string, error) {
	res, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}
