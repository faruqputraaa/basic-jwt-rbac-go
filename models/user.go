package models

import "time"

// Model untuk User
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100)"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"type:varchar(20);not null;default:'Editor'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
