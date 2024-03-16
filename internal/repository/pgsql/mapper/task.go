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

func ToDomainTask(mdl *models.Task) *entity.Task {
	return &entity.Task{
		ID:          mdl.ID,
		UserId:      mdl.UserID,
		Title:       mdl.Title,
		Description: mdl.Description,
		Status:      mdl.Status,
		CreatedAt:   mdl.CreatedAt,
		UpdatedAt:   mdl.UpdatedAt,
	}
}

func ToListDomainTask(models []*models.Task) []*entity.Task {
	var tasks []*entity.Task
	for _, mdl := range models {
		tasks = append(tasks, ToDomainTask(mdl))
	}
	return tasks
}
