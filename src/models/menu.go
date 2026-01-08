package models

import "time"

type Menu struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"type:varchar(255);not null"`
	BasePrice   float32   `gorm:"not null"`
	Description string    `gorm:"type:text"`
	Image       string    `gorm:"type:varchar(255)"`
	Products    []Product `gorm:"many2many:menu_products;constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}
