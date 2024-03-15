package repository

import (
	"context"
	"fun-with-gin/domain/entity"
)

type UserRepository interface {
	FindOne(ctx context.Context, filter *entity.UserFilter) (*entity.User, error)
}
