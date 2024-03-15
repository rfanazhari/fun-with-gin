package pgsql_user

import (
	"context"
	"errors"
	"fun-with-gin/domain/entity"
	"fun-with-gin/internal/repository/pgsql/mapper"
	"fun-with-gin/internal/repository/pgsql/models"
)

func (u userPgsqlInit) FindOne(ctx context.Context, filter *entity.UserFilter) (*entity.User, error) {
	var mdl *models.User

	tx := u.db.WithContext(ctx)

	if filter.Email != "" && filter.ID != 0 {
		tx = tx.Where("email = ?", filter.Email).Where("id = ?", filter.ID)
	} else if filter.Email != "" {
		tx = tx.Where("email = ?", filter.Email)
	} else if filter.ID != 0 {
		tx = tx.Where("id = ?", filter.ID)
	} else {
		return nil, errors.New("empty filter")
	}

	if err := tx.First(mdl).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}
		return nil, err
	}

	return mapper.ToDomainUser(mdl), nil
}
