package service

import (
	"wackdo/src/models"
	"wackdo/src/service/repository"
)

func GetProductById(id int) (models.Product, error) {
	return repository.GetSellableItemByID[models.Product](id)
}

func GetProductByName(name string) (models.Product, error) {
	return repository.GetSellableItemByName[models.Product](name)
}

func GetProducts(page, pageSize int) ([]models.Product, error) {
	return repository.GetAllSellableItems[models.Product](page, pageSize)
}

func DeleteProductById(id int) error {
	return repository.DeleteSellableItemByID[models.Product](id)
}

func ProductExists(name string) (bool, error) {
	count, err := repository.GetAllSellableItemByName[models.Product](name)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func CreateProduct(product models.Product) (models.Product, error) {
	created, err := repository.CreateSellableItem(&product)
	if err != nil {
		return models.Product{}, err
	}
	return *created, nil
}
