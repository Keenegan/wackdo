package models

type Order struct {
	ID     uint
	Status string
}

type OrderLine struct {
	ID        uint
	unitPrice float32
	item      SellableItem
}
