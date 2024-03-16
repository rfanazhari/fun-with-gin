package login_token

import (
	"fun-with-gin/domain/service"
	pkgSvc "fun-with-gin/pkg/domain/service"
)

type redisLoginToken struct {
	client pkgSvc.RedisClient
}

func NewRedisLoginToken(client pkgSvc.RedisClient) service.LoginTokenService {
	return &redisLoginToken{
		client: client,
	}
}
