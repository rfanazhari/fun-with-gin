package pgsql_task

import (
	"fun-with-gin/domain/repository"
	"gorm.io/gorm"
)

type taskPgsqlInit struct {
	db *gorm.DB
}

func NewTaskPgsqlInit(db *gorm.DB) repository.TaskRepository {
	return taskPgsqlInit{db: db}
}
