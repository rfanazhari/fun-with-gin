package usecase_task

import (
	"context"
	"fun-with-gin/domain/entity"
)

func (u taskUseCaseInit) ListTask(ctx context.Context, filter *entity.TaskFilter) ([]*entity.Task, error) {
	tasks, err := u.taskRepo.FindList(ctx, filter)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
