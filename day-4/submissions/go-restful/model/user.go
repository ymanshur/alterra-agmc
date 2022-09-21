package model

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Name      string    `gorm:"size:256" json:"name" validate:"required"`
	Email     string    `gorm:"not null;unique" json:"email" validate:"required,email"`
	Password  string    `gorm:"not null" json:"password" validate:"required"`
	Token     string    `gorm:"-:all" json:"token"`
}
