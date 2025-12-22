package models

import (
	"gorm.io/datatypes"
)

type Employee struct {
	ID    uint
	Name  string
	Roles datatypes.JSONSlice[Role]
}

type Order struct {
	ID     uint
	Status string
}

type OrderLine struct {
	ID        uint
	unitPrice float32
	item      SellableItem
}
