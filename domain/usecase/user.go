package usecase

import (
	"context"
	"fun-with-gin/domain/entity"
)

type UserUseCase interface {
	CreateOne(ctx context.Context, payload *entity.User) error
	FindOne(ctx context.Context, filter *entity.UserFilter) (*entity.User, error)
	UpdateOne(ctx context.Context, payload *entity.User, filter *entity.UserFilter) error
	ListUser(ctx context.Context, filter *entity.UserFilter) ([]*entity.User, error)
}
