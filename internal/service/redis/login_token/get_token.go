package login_token

import (
	"context"
	"fun-with-gin/domain/service"
)

func (s redisLoginToken) GetToken(ctx context.Context, userId string) (string, error) {
	key := s.key(service.RedisKeyStoreToken, userId)
	return s.client.Get(ctx, key)
}
