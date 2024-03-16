package usecase_user

import (
	"context"
	"fun-with-gin/domain/repository"
	"fun-with-gin/domain/service"
	"fun-with-gin/domain/usecase"
)

type userUseCaseInit struct {
	ctx           context.Context
	userRepo      repository.UserRepository
	storeTokenSvc service.LoginTokenService
}

func NewUseCaseUser(ctx context.Context, userRepo repository.UserRepository, storeTknSvc service.LoginTokenService) usecase.UserUseCase {
	return &userUseCaseInit{
		ctx:           ctx,
		userRepo:      userRepo,
		storeTokenSvc: storeTknSvc,
	}
}
