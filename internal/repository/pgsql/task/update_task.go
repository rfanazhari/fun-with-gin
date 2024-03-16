package pgsql_task

import (
	"context"
	"fun-with-gin/domain/entity"
	"fun-with-gin/internal/repository/pgsql/mapper"
)

func (t taskPgsqlInit) UpdateTask(ctx context.Context, payload entity.Task) error {
	tx := t.db.WithContext(ctx)
	mdl := mapper.ToModelTask(payload)

	query := tx.Model(&mdl)

	if err := query.Updates(&mdl).Error; err != nil {
		return err
	}

	return nil
}
