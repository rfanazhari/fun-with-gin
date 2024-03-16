package usecase

import (
	"context"
	"fun-with-gin/domain/entity"
)

type TaskUseCase interface {
	CreateOne(ctx context.Context, payload *entity.Task) error
	FindOne(ctx context.Context, filter *entity.TaskFilter) (*entity.Task, error)
	UpdateOne(ctx context.Context, payload *entity.Task, filter *entity.TaskFilter) error
	ListTask(ctx context.Context, filter *entity.TaskFilter) ([]*entity.Task, error)
}
