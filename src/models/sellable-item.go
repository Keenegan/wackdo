package models

type SellableItem struct {
	ID uint
}

type Product struct {
	ID          uint
	Name        string
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
	Name        string
	BasePrice   float32
	Description string
	Image       string
	Products    []Product `gorm:"many2many:menu_products;constraint:OnDelete:CASCADE;"`
}
