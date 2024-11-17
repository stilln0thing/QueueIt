package models


type Column struct {
	ID         uint         `gorm:"primaryKey"`
	Title      string       `gorm:"not null"`
	MaxLimit   int          `gorm:"not null"`
	EventID    uint         `gorm:"not null"`
    Customers  []Customer   `gorm:"foreignKey:ColumnID"`     // one queuecolumn can have many customers 


	Event     Event         `gorm:"foreignKey:EventID"`
}