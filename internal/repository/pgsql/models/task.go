package models

import "time"

type Task struct {
	ID          uint      `gorm:"primaryKey,column:id"`
	UserID      uint      `gorm:"column:user_id"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
	Status      string    `gorm:"column:status"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP,column:created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP,column:updated_at"`
}

func (Task) TableName() string {
	return "tasks"
}
