package repository

import (
	"context"
	"fun-with-gin/domain/entity"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, payload entity.Task) error
	FindList(ctx context.Context, filter *entity.TaskFilter) ([]*entity.Task, error)
	FindOneTask(ctx context.Context, filter *entity.TaskFilter) (*entity.Task, error)
}
