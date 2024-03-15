package usecase_user

import (
	"context"
	"errors"
	"fun-with-gin/domain/entity"
)

func (u userUseCaseInit) FindOne(ctx context.Context, filter *entity.UserFilter) (*entity.User, error) {
	user, err := u.userRepo.FindOneUser(ctx, &entity.UserFilter{ID: filter.ID})
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
