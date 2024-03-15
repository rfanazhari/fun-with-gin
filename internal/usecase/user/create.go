package usecase_user

import (
	"context"
	"errors"
	"fmt"
	"fun-with-gin/domain/entity"
)

func (u userUseCaseInit) CreateOne(ctx context.Context, payload *entity.User) error {
	user, err := u.userRepo.FindOneUser(ctx, &entity.UserFilter{Email: &payload.Email})
	if err != nil {
		return err
	}
	if user != nil {
		return errors.New("user already exists")
	}
	fmt.Println(payload)
	insert := u.userRepo.InsertUser(ctx, *payload)
	if insert != nil {
		return insert
	}
	return nil
}
