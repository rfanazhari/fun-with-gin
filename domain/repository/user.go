package repository

import (
	"context"
	"fun-with-gin/domain/entity"
)

type UserRepository interface {
	FindOneUser(ctx context.Context, filter *entity.UserFilter) (*entity.User, error)
}
