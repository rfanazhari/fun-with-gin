package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey,column:id"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"unique,column:email"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP,column:created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP,column:updated_at"`
}

func (User) TableName() string {
	return "users"
}
