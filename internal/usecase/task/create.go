package usecase_task

import (
	"context"
	"fun-with-gin/domain/entity"
)

func (u taskUseCaseInit) CreateOne(ctx context.Context, payload *entity.Task) error {
	err := u.taskRepo.CreateTask(ctx, *payload)
	if err != nil {
		return err
	}
	return nil
}
