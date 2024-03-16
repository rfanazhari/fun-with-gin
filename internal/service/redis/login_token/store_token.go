package login_token

import (
	"context"
	"fun-with-gin/domain/service"
	"time"
)

func (s redisLoginToken) StoreToken(ctx context.Context, userId string, email any, exp time.Duration) error {
	key := s.key(service.RedisKeyStoreToken, userId)
	return s.client.Set(ctx, key, email, exp)
}
