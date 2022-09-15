package models

import (
	"time"
)

type User struct {
	// gorm.Model
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
	Name      string     `json:"name" form:"name"`
	Email     string     `json:"email" form:"email"`
	Password  string     `json:"password" form:"password"`
}
