package repository

import (
	"context"
	"fun-with-gin/domain/entity"
)

type UserRepository interface {
	FindOneUser(ctx context.Context, filter *entity.UserFilter) (*entity.User, error)
	FindList(ctx context.Context) ([]*entity.User, error)
	InsertUser(ctx context.Context, payload entity.User) error
}
