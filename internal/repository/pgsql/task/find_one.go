package pgsql_task

import (
	"context"
	"fun-with-gin/domain/entity"
	"fun-with-gin/internal/repository/pgsql/mapper"
	"fun-with-gin/internal/repository/pgsql/models"
)

func (t taskPgsqlInit) FindOneTask(ctx context.Context, filter *entity.TaskFilter) (*entity.Task, error) {
	var taskList *models.Task

	tx := t.db.WithContext(ctx)

	if filter != nil {
		if filter.ID != 0 {
			tx = tx.Where("id = ?", filter.ID)
		}
		if filter.UserId != 0 {
			tx = tx.Where("user_id = ?", filter.UserId)
		}
	}

	if err := tx.First(&taskList).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}
		return nil, err
	}

	return mapper.ToDomainTask(taskList), nil
}
