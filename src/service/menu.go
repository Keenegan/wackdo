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

func MenuExists(name string) (bool, error) {
	count, err := repository.GetAllSellableItemByName[models.Menu](name)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func CreateMenu(menu models.Menu) (models.Menu, error) {
	created, err := repository.CreateSellableItem(&menu)
	if err != nil {
		return models.Menu{}, err
	}
	return *created, nil
}
