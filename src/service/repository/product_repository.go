package product_repository

import (
	"errors"
	"wackdo/src/initializers"
	"wackdo/src/models"
)

func GetProductById(id int) (models.Product, error) {
	var product models.Product

	if id <= 0 {
		return product, errors.New("invalid product id")
	}

	err := initializers.DB.First(&product, id).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func GetProductByName(name string) (models.Product, error) {
	var product models.Product

	err := initializers.DB.Where("name = ?", name).First(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func GetProducts(page, pageSize int) ([]models.Product, error) {
	var products []models.Product

	err := initializers.DB.Order("id ASC").Limit(pageSize).Offset(page).Find(&products).Error
	if err != nil {
		return products, err
	}
	return products, err

}
