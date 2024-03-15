package usecase_user

import (
	"context"
	"fun-with-gin/domain/entity"
)

func (u userUseCaseInit) ListUser(ctx context.Context) ([]*entity.User, error) {
	users, err := u.userRepo.FindList(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
