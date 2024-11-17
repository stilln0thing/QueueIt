package models

import (
	"github.com/google/uuid"
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID            uint           `gorm:"primaryKey;autoIncrement"` 
	FirstName     string         `gorm:"size:100;not null" validate:"required"`
	LastName      string         `gorm:"size:100;not null" validate:"required"`
	Password      string         `gorm:"not null" validate:"required"`
	Email         string         `gorm:"size:100;unique;not null" validate:"required,email"`
	Phone         string         `gorm:"size:15;not null" validate:"required"`
	Token         string         `gorm:"size:255"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime"`
	CustomerID    string         `gorm:"size:36;unique"`
	
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    if u.CustomerID == "" {
        u.CustomerID = uuid.New().String() 
    }
    return
}