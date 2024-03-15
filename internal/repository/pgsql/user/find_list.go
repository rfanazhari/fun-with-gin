package pgsql_user

import (
	"context"
	"fun-with-gin/domain/entity"
	"fun-with-gin/internal/repository/pgsql/mapper"
	"fun-with-gin/internal/repository/pgsql/models"
)

func (u userPgsqlInit) FindList(ctx context.Context) ([]*entity.User, error) {
	var mdls []*models.User

	tx := u.db.WithContext(ctx)

	if err := tx.Find(&mdls).Error; err != nil {
		return nil, err
	}

	return mapper.ToListDomainUser(mdls), nil
}
