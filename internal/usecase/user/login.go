package usecase_user

import (
	"context"
	"errors"
	"fun-with-gin/domain/entity"
	"fun-with-gin/pkg/utils"
)

func (u *userUseCaseInit) LoginUser(ctx context.Context, email, password string) (string, error) {
	user, err := u.userRepo.FindOneUser(ctx, &entity.UserFilter{Email: &email})
	if err != nil {
		return "", err
	}

	validPassword := utils.ValidatePassword(user.Password, password)
	if !validPassword {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateToken(utils.JwtPayload{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		return "", err
	}

	return token, nil
}
