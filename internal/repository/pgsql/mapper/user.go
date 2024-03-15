package mapper

import (
	"fun-with-gin/domain/entity"
	"fun-with-gin/internal/repository/pgsql/models"
)

func ToDomainUser(m *models.User) *entity.User {
	return &entity.User{
		ID:        m.ID,
		Name:      m.Name,
		Email:     m.Email,
		Password:  m.Password,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
