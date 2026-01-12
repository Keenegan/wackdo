package service

import (
	"errors"
	"wackdo/src/initializers"
	"wackdo/src/models"

	"gorm.io/gorm"
)

func GetProductById(id int) (models.Product, error) {
	var product models.Product

	if id <= 0 {
		return product, &EntityNotFoundError{models.Product{}}
	}

	err := initializers.DB.First(&product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return product, &EntityNotFoundError{models.Product{}}
		}
		return product, err
	}

	return product, nil
}

func GetProductByName(name string) (models.Product, error) {
	var product models.Product

	err := initializers.DB.
		Where("name = ?", name).
		First(&product).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return product, &EntityNotFoundError{models.Product{}}
		}
		return product, err
	}

	return product, nil
}

func GetProducts(page, pageSize int) ([]models.Product, error) {
	var products []models.Product

	err := initializers.DB.
		Order("id ASC").
		Limit(pageSize).
		Offset(page).
		Find(&products).
		Error

	if err != nil {
		return products, err
	}

	return products, nil
}

func DeleteProductById(id int) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	return initializers.DB.Transaction(func(tx *gorm.DB) error {
		// Find all menus containing this product
		var menus []models.Menu
		err := tx.
			Joins("JOIN menu_products ON menu_products.menu_id = menus.id").
			Where("menu_products.product_id = ?", id).
			Find(&menus).Error
		if err != nil {
			return err
		}

		// Delete all menus that contain this product
		for _, menu := range menus {
			if err := tx.Delete(&menu).Error; err != nil {
				return err
			}
		}

		// Delete the product
		return tx.Delete(&models.Product{}, id).Error
	})
}

func ProductExists(name string) (bool, error) {
	var count int64

	err := initializers.DB.
		Model(&models.Product{}).
		Where("name = ?", name).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func GetProductsByIds(ids []uint) ([]models.Product, error) {
	var products []models.Product

	err := initializers.DB.
		Model(&models.Product{}).
		Where("id IN ?", ids).
		Find(&products).
		Error

	if err != nil {
		return products, err
	}

	return products, nil
}

func CreateProduct(product models.Product) (models.Product, error) {
	if err := initializers.DB.Create(&product).Error; err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func UpdateProduct(product models.Product) (models.Product, error) {
	if err := initializers.DB.Save(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}
