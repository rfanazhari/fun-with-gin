package responses

import (
	"fun-with-gin/domain/entity"
	"time"
)

type UserResponse struct {
	ID        uint      `json:"ID"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ToUserResponses(usr *entity.User) *UserResponse {
	return &UserResponse{
		ID:        usr.ID,
		Name:      usr.Name,
		Email:     usr.Email,
		CreatedAt: usr.CreatedAt,
		UpdatedAt: usr.UpdatedAt,
	}
}

func ToListUserResponses(usr []*entity.User) []*UserResponse {
	var users []*UserResponse
	for _, u := range usr {
		users = append(users, ToUserResponses(u))
	}
	return users
}
