package models

import "time"

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusConfirmed OrderStatus = "confirmed"
	OrderStatusPreparing OrderStatus = "preparing"
	OrderStatusReady     OrderStatus = "ready"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

func (s OrderStatus) IsValid() bool {
	switch s {
	case OrderStatusPending, OrderStatusConfirmed, OrderStatusPreparing, OrderStatusReady, OrderStatusDelivered, OrderStatusCancelled:
		return true
	default:
		return false
	}
}

type Order struct {
	ID         uint        `gorm:"primaryKey"`
	Status     OrderStatus `gorm:"type:varchar(20);not null;default:'pending'"`
	TotalPrice float32     `gorm:"not null"`
	OrderLines []OrderLine `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	CreatedAt  time.Time   `gorm:"not null"`
	UpdatedAt  time.Time   `gorm:"not null"`
}

type ItemType string

const (
	ItemTypeProduct ItemType = "product"
	ItemTypeMenu    ItemType = "menu"
)

type OrderLine struct {
	ID          uint     `gorm:"primaryKey"`
	OrderID     uint     `gorm:"not null;index"`
	ItemType    ItemType `gorm:"type:varchar(10);not null"`
	ItemName    string   `gorm:"not null"`
	UnitPrice   float32  `gorm:"not null"`
	Quantity    int      `gorm:"not null;default:1"`
	Description string   `gorm:"type:text"`
	Content		string	 `gorm:"type:text"`
	Image       string   `gorm:"type:varchar(255)"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}
