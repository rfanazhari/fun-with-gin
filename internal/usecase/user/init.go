package usecase_user

import (
	"context"
	"fun-with-gin/domain/repository"
	"fun-with-gin/domain/usecase"
)

type userUseCaseInit struct {
	ctx      context.Context
	userRepo repository.UserRepository
}

func NewUseCaseUser(ctx context.Context) usecase.UserUseCase {
	return &userUseCaseInit{ctx: ctx}
}
