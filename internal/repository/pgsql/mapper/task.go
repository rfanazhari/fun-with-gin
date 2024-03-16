package mapper

import (
	"fun-with-gin/domain/entity"
	"fun-with-gin/internal/repository/pgsql/models"
)

func ToModelTask(d entity.Task) *models.Task {
	return &models.Task{
		ID:          d.ID,
		UserID:      d.UserId,
		Title:       d.Title,
		Description: d.Description,
		Status:      d.Status,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}
