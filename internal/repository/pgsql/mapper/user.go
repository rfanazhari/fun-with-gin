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
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ToListDomainUser(m []*models.User) []*entity.User {
	var users []*entity.User
	for _, u := range m {
		users = append(users, ToDomainUser(u))
	}

	return users
}

func ToModelUser(m *entity.User) *models.User {
	return &models.User{
		ID:        m.ID,
		Name:      m.Name,
		Email:     m.Email,
		Password:  m.Password,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
