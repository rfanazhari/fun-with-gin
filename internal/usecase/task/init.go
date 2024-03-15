package usecase_task

import (
	"context"
	"fun-with-gin/domain/usecase"
)

type taskUseCaseInit struct {
	ctx context.Context
}

func NewUseCaseTask(ctx context.Context) usecase.TaskUseCase {
	return &taskUseCaseInit{
		ctx: ctx,
	}
}
