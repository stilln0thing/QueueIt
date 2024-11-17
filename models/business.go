package models

import(
	// "github.com/google/uuid"
	// "gorm.io/gorm"
)

type Business struct {
	ID                      uint 	       `gorm:"primaryKey, autoIncrement"`
	BusinessName 	        string 	       `gorm:"not null"`
	BusinessAddress	        string 	       `gorm:"not null"`
	BusinessCategory        Category       `gorm:"not null"`
	BusinessCurrentEvents   []Event        `gorm:"foreignKey:BusinessID"` 
	// LastClosedEvent    //will implement later         
}

// func (u *Business) BeforeCreate(tx *gorm.DB) (err error) {
//     if u.BusinessID == "" {
//         u.BusinessID = uuid.New().String() 
//     }
//     return
// }

func (Business) TableName() string {
	return "Business" // Use singular "business" instead of "businesses"
}