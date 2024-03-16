package entity

import (
	"fun-with-gin/pkg/utils"
	"time"
)

type Task struct {
	ID          uint
	UserId      uint
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TaskDto struct {
	UserId      uint
	Title       string
	Description string
	Status      string
}

type TaskFilter struct {
	ID     uint
	UserId uint
}

func NewTask(payload TaskDto) *Task {
	id := utils.GenerateID()
	user := Task{
		ID:          uint(id),
		UserId:      payload.UserId,
		Title:       payload.Title,
		Description: payload.Description,
		Status:      payload.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return &user
}
