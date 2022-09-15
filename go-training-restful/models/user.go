package models

import (
	"time"
)

type User struct {
	// gorm.Model
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	Name      *string   `gorm:"size:256" json:"name" form:"name"`
	Email     string    `gorm:"not null;unique" json:"email" form:"email"`
	Password  string    `gorm:"not null" json:"password" form:"password"`
	Token     string    `gorm:"-:all" json:"token"`
}
