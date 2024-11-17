package models

import (
	"time"
)



type Event struct{
	ID       uint           `gorm:"primaryKey;autoIncrement"`
	BusinessID    uint           `gorm:"not null"`
	OpeningTime   time.Time		 `gorm:"not null"`
	ClosingTime   time.Time		 `gorm:"not null"`
	Columns       []Column       `gorm:"foreignKey:EventID"`   // one event can have many queueColumns 
	Category	  Category       `gorm:"not null"`
	Title         string         `gorm:"not null"`
	Description   string         `gorm:"size:200"`
    AvgWaitTime   int 	         `gorm:"not null"`

	Business     Business `gorm:"foreignKey:BusinessID"`
}