package pgsql_user

import (
	"context"
	"errors"
	"fun-with-gin/domain/entity"
	"fun-with-gin/internal/repository/pgsql/mapper"
	"fun-with-gin/internal/repository/pgsql/models"
)

func (u userPgsqlInit) FindOneUser(ctx context.Context, filter *entity.UserFilter) (*entity.User, error) {
	var mdl models.User

	tx := u.db.WithContext(ctx)

	if filter != nil {
		if filter.ID != nil && filter.Email != nil {
			tx = tx.Where("email = ?", filter.Email).Where("id = ?", filter.ID)
		} else if filter.ID != nil {
			tx = tx.Where("id = ?", *filter.ID)
		} else if filter.Email != nil {
			tx = tx.Where("email = ?", *filter.Email)
		}
	} else {
		return nil, errors.New("empty filter")
	}

	if err := tx.First(&mdl).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}
		return nil, err
	}

	return mapper.ToDomainUser(&mdl), nil
}
