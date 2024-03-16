package service

import (
	"context"
	"time"
)

type RedisKey string

const (
	RedisKeyStoreToken RedisKey = "token_login"
)

type LoginTokenService interface {
	StoreToken(ctx context.Context, key string, value any, exp time.Duration) error
	GetToken(ctx context.Context, key string) (string, error)
	RevokeToken(ctx context.Context, key string) error
}
