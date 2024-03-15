package usecase_user

import (
	"context"
	"fun-with-gin/domain/entity"
	"fun-with-gin/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

func (u userUseCaseInit) LoginUser(ctx context.Context, email, password string) (string, error) {
	user, err := u.userRepo.FindOne(ctx, &entity.UserFilter{Email: email})
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
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
