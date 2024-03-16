package repository

import (
	"context"
	"fun-with-gin/domain/entity"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, payload entity.Task) error
}
