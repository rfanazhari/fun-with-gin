package usecase_task

import (
	"context"
	"fmt"
	"fun-with-gin/domain/entity"
)

func (u taskUseCaseInit) FindOne(ctx context.Context, filter *entity.TaskFilter) (*entity.Task, error) {
	task, err := u.taskRepo.FindOneTask(ctx, filter)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, fmt.Errorf("task not found")
	}
	return task, nil
}
