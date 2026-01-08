package service

import (
	"errors"
	"strings"
	"wackdo/src/initializers"
	"wackdo/src/models"

	"gorm.io/gorm"
)

type OrderItemRequest struct {
	ItemType ItemType `json:"itemType"`
	ItemID   uint     `json:"itemId"`
	Quantity int      `json:"quantity"`
}

type ItemType string

const (
	ItemTypeProduct ItemType = "product"
	ItemTypeMenu    ItemType = "menu"
)

func CreateOrder(items []OrderItemRequest) (models.Order, error) {
	if len(items) == 0 {
		return models.Order{}, errors.New("order must contain at least one item")
	}

	var orderLines []models.OrderLine
	var totalPrice float32 = 0

	for _, item := range items {
		if item.Quantity <= 0 {
			return models.Order{}, errors.New("quantity must be greater than 0")
		}

		orderLine, err := createOrderLineFromItem(item)
		if err != nil {
			return models.Order{}, err
		}

		totalPrice += orderLine.UnitPrice * float32(orderLine.Quantity)
		orderLines = append(orderLines, orderLine)
	}

	order := models.Order{
		Status:     models.OrderStatusPending,
		TotalPrice: totalPrice,
		OrderLines: orderLines,
	}

	if err := initializers.DB.Create(&order).Error; err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func createOrderLineFromItem(item OrderItemRequest) (models.OrderLine, error) {
	switch item.ItemType {
	case ItemTypeProduct:
		return createOrderLineFromProduct(item)
	case ItemTypeMenu:
		return createOrderLineFromMenu(item)
	default:
		return models.OrderLine{}, errors.New("invalid item type")
	}
}

func createOrderLineFromProduct(item OrderItemRequest) (models.OrderLine, error) {
	product, err := GetProductById(int(item.ItemID))
	if err != nil {
		return models.OrderLine{}, err
	}

	if !product.Available {
		return models.OrderLine{}, errors.New("product is not available")
	}

	return models.OrderLine{
		ItemType:    models.ItemTypeProduct,
		ItemName:    product.Name,
		UnitPrice:   product.BasePrice,
		Quantity:    item.Quantity,
		Description: product.Description,
		Image:       product.Image,
	}, nil
}

func createOrderLineFromMenu(item OrderItemRequest) (models.OrderLine, error) {
	menu, err := GetMenuById(int(item.ItemID))
	if err != nil {
		return models.OrderLine{}, err
	}

	description := []string{}
	for _, v := range menu.Products {
		description = append(description, v.Name)
	}

	return models.OrderLine{
		ItemType:    models.ItemTypeMenu,
		ItemName:    menu.Name,
		UnitPrice:   menu.BasePrice,
		Quantity:    item.Quantity,
		Description: menu.Description,
		Content:     strings.Join(description, " + "),
		Image:       menu.Image,
	}, nil
}

func GetOrders(page, pageSize int) ([]models.Order, error) {
	var orders []models.Order

	err := initializers.DB.
		Preload("OrderLines").
		Order("created_at DESC").
		Limit(pageSize).
		Offset(page * pageSize).
		Find(&orders).
		Error

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func GetOrderById(id uint) (models.Order, error) {
	var order models.Order

	err := initializers.DB.
		Preload("OrderLines").
		First(&order, id).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Order{}, &EntityNotFoundError{models.Order{}}
		}
		return models.Order{}, err
	}

	return order, nil
}

func UpdateOrderStatus(id uint, status models.OrderStatus) (models.Order, error) {
	if !status.IsValid() {
		return models.Order{}, errors.New("invalid order status")
	}

	var order models.Order
	if err := initializers.DB.First(&order, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Order{}, &EntityNotFoundError{models.Order{}}
		}
		return models.Order{}, err
	}

	order.Status = status
	if err := initializers.DB.Save(&order).Error; err != nil {
		return models.Order{}, err
	}

	return order, nil
}
