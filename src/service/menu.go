package service

import (
	"wackdo/src/models"
	"wackdo/src/service/repository"
)

func GetMenuById(id int) (models.Menu, error) {
	return repository.GetSellableItemByID[models.Menu](id)
}

func GetMenuByName(name string) (models.Menu, error) {
	return repository.GetSellableItemByName[models.Menu](name)
}

func GetMenus(page, pageSize int) ([]models.Menu, error) {
	return repository.GetAllSellableItems[models.Menu](page, pageSize)
}

func DeleteMenuById(id int) error {
	return repository.DeleteSellableItemByID[models.Menu](id)
}
