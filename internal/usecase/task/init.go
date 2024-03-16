package usecase_task

import (
	"fun-with-gin/domain/repository"
	"fun-with-gin/domain/usecase"
)

type taskUseCaseInit struct {
	taskRepo repository.TaskRepository
}

func NewUseCaseTask(taskRepo repository.TaskRepository) usecase.TaskUseCase {
	return &taskUseCaseInit{
		taskRepo: taskRepo,
	}
}
