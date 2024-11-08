package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID            uint           `gorm:"primaryKey;autoIncrement"` 
	FirstName     string         `json:"first_name" gorm:"size:100;not null" validate:"required"`
	LastName      string         `json:"last_name" gorm:"size:100;not null" validate:"required"`
	Password      string         `json:"password" gorm:"not null" validate:"required"`
	Email         string         `json:"email" gorm:"size:100;unique;not null" validate:"required,email"`
	Phone         string         `json:"phone" gorm:"size:15;not null" validate:"required"`
	Token         string         `json:"token" gorm:"size:255"`
	UserType      string         `json:"user_type" gorm:"size:50;not null" validate:"required"`
	RefreshToken  string         `json:"refresh_token" gorm:"size:255"`
	CreatedAt     time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	UserID        string         `json:"user_id" gorm:"size:100;unique;not null" validate:"required,uuid"`
	DeletedAt     gorm.DeletedAt `gorm:"index"` // Soft delete support 
}