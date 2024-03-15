package usecase_user

import (
	"context"
	"errors"
	"fun-with-gin/domain/entity"
)

func (u userUseCaseInit) UpdateOne(ctx context.Context, payload *entity.User, filter *entity.UserFilter) error {
	user, err := u.userRepo.FindOneUser(ctx, filter)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	if payload.Email != "" {
		existingUser, err := u.userRepo.FindOneUser(ctx, &entity.UserFilter{Email: &payload.Email})
		if err != nil {
			return err
		}
		if existingUser != nil && existingUser.ID != user.ID {
			return errors.New("email already exists")
		}
	}

	payload.ID = user.ID
	err = u.userRepo.UpdateUser(ctx, *payload, filter)
	if err != nil {
		return err
	}
	return nil
}
