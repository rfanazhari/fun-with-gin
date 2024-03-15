package pgsql_user

import (
	"context"
	"fun-with-gin/domain/entity"
	"fun-with-gin/internal/repository/pgsql/mapper"
)

func (u userPgsqlInit) InsertUser(ctx context.Context, payload entity.User) error {
	tx := u.db.WithContext(ctx)
	mdl := mapper.ToModelUser(&payload)

	if err := tx.Create(mdl).Error; err != nil {
		return err
	}

	return nil
}
