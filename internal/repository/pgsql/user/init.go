package pgsql_user

import (
	"fun-with-gin/domain/repository"
	"gorm.io/gorm"
)

type userPgsqlInit struct {
	db *gorm.DB
}

func NewUserPgsqlInit(db *gorm.DB) repository.UserRepository {
	return &userPgsqlInit{
		db: db,
	}
}
