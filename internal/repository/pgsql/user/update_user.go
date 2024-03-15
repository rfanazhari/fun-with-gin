package pgsql_user

import (
	"context"
	"fun-with-gin/domain/entity"
	"fun-with-gin/internal/repository/pgsql/mapper"
)

func (u userPgsqlInit) UpdateUser(ctx context.Context, payload entity.User, filter *entity.UserFilter) error {
	tx := u.db.WithContext(ctx)
	mdl := mapper.ToModelUser(&payload)

	query := tx.Model(&mdl)

	if filter != nil {
		if filter.ID != nil {
			query = query.Where("id = ?", filter.ID)
		}
	}

	if err := query.Updates(&mdl).Error; err != nil {
		return err
	}

	return nil
}
