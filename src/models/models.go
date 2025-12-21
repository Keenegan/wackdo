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

type SellableItem struct {
	ID          uint
	name        string
	basePrice   float32
	description string
	image       string
}

type Product struct {
	ID       uint
	category string
}

type Menu struct {
	ID uint
}

type MenuOption struct {
	ID    uint
	price float32
	name  string
}
