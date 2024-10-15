package models

import (
	"time"
)

// Model untuk Todo
type Todo struct {
	ID        	uint      `gorm:"primaryKey"`
	Title     	string    `gorm:"type:varchar(100);not null"`
	Content   	string    `gorm:"type:text;not null"`
	UserID    	uint      `json:"user_id"`
	Completed	bool      `gorm:"default:false"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

