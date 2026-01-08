package models

import "time"

type Category string

const (
	CategoryFood  Category = "food"
	CategoryDrink Category = "drink"
)

func (c Category) IsValid() bool {
	switch c {
	case CategoryDrink, CategoryFood:
		return true
	default:
		return false
	}
}

type Product struct {
	ID          uint     `gorm:"primaryKey"`
	Name        string   `gorm:"type:varchar(255);not null"`
	BasePrice   float32  `gorm:"not null"`
	Description string   `gorm:"type:text"`
	Image       string   `gorm:"type:varchar(255)"`
	Category    Category `gorm:"type:varchar(20);not null"`
	Available   bool     `gorm:"not null;default:true"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}
