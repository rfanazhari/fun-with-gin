package usecase_task

import (
	"context"
	"fmt"
	"fun-with-gin/domain/entity"
	"time"
)

func (u taskUseCaseInit) UpdateOne(ctx context.Context, payload *entity.Task, filter *entity.TaskFilter) error {
	task, err := u.taskRepo.FindOneTask(ctx, filter)
	if err != nil {
		return err
	}
	if task == nil {
		return fmt.Errorf("task not found")
	}
	task.Title = payload.Title
	task.Description = payload.Description
	task.Status = payload.Status
	task.UpdatedAt = time.Now()

	err = u.taskRepo.UpdateTask(ctx, *task)
	if err != nil {
		return err
	}

	return nil
}
