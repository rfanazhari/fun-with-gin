package login_token

import (
	"context"
	"fun-with-gin/domain/service"
)

func (s redisLoginToken) RevokeToken(ctx context.Context, userId string) error {
	key := s.key(service.RedisKeyStoreToken, userId)
	return s.client.Del(ctx, key)
}
