package service

import (
	"errors"
	"wackdo/src/initializers"
	"wackdo/src/models"

	"gorm.io/gorm"
)

func GetMenuById(id int) (models.Menu, error) {
	var menu models.Menu

	if id <= 0 {
		return menu, &EntityNotFoundError{models.Menu{}}
	}

	err := initializers.DB.
		Preload("Products").
		First(&menu, id).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return menu, &EntityNotFoundError{models.Menu{}}
		}
		return menu, err
	}

	return menu, nil
}

func GetMenuByName(name string) (models.Menu, error) {
	var menu models.Menu

	err := initializers.DB.
		Preload("Products").
		Where("name = ?", name).
		First(&menu).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return menu, &EntityNotFoundError{models.Menu{}}
		}
		return menu, err
	}

	return menu, nil
}

func GetMenus(page, pageSize int) ([]models.Menu, error) {
	var menus []models.Menu

	err := initializers.DB.
		Preload("Products").
		Order("id ASC").
		Limit(pageSize).
		Offset(page).
		Find(&menus).
		Error

	if err != nil {
		return menus, err
	}

	return menus, nil
}

func DeleteMenuById(id int) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	return initializers.DB.Delete(&models.Menu{}, id).Error
}

func MenuExists(name string) (bool, error) {
	var count int64

	err := initializers.DB.
		Model(&models.Menu{}).
		Where("name = ?", name).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func CreateMenu(menu models.Menu) (models.Menu, error) {
	if err := initializers.DB.Create(&menu).Error; err != nil {
		return models.Menu{}, err
	}
	return menu, nil
}

func UpdateMenu(menu models.Menu) (models.Menu, error) {
	if err := initializers.DB.Save(&menu).Error; err != nil {
		return menu, err
	}
	return menu, nil
}
