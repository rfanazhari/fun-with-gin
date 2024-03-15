package entity

import (
	"fun-with-gin/pkg/utils"
	"time"
)

type User struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserFilter struct {
	ID    uint
	Email *string
}

type UserDto struct {
	Name     string
	Email    string
	Password string
}

func NewUser(payload UserDto) *User {
	id := utils.GenerateID()
	user := &User{
		ID:        uint(id),
		Name:      payload.Name,
		Email:     payload.Email,
		Password:  payload.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return user
}
