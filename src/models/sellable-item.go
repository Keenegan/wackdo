package models

type SellableItem struct {
	ID uint
}

type Product struct {
	ID          uint
	Name        string `gorm:"uniqueIndex"`
	BasePrice   float32
	Description string
	Image       string
	Category    Category
	Available   bool
}

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

type Menu struct {
	ID          uint
	name        string `gorm:"uniqueIndex"`
	basePrice   float32
	description string
	image       string
	options     []MenuOption
}

type MenuOption struct {
	ID    uint
	price float32
	name  string
}
