package repository

import (
	"errors"
	"wackdo/src/initializers"
	"wackdo/src/models"
)

// todo see if we can use interface instead of any
func GetSellableItemByID[T any](id int) (T, error) {
	var entity T

	if id <= 0 {
		return entity, errors.New("invalid id")
	}

	err := initializers.DB.First(&entity, id).Error
	if err != nil {
		return entity, err
	}

	return entity, nil
}

func GetSellableItemByName[T any](name string) (T, error) {
	var entity T

	db := initializers.DB
	
	switch any(&entity).(type) {
	case *models.Menu:
		db = db.Preload("Products")
	}

	err := db.
		Where("name = ?", name).
		First(&entity).
		Error

	return entity, err
}

// todo rename
func GetAllSellableItemByName[T any](name string) (int64, error) {
	var entity T
	var count int64

	err := initializers.DB.
		Model(&entity).
		Where("name = ?", name).
		Count(&count).
		Error

	if err != nil {
		return count, err
	}

	return count, nil
}

func GetAllSellableItemById[T any](ids []uint) ([]T, error) {
	var entity T
	var results []T

	err := initializers.DB.
		Model(&entity).
		Where("id IN ?", ids).
		Find(&results).
		Error

	if err != nil {
		return results, err
	}

	return results, nil
}

func GetAllSellableItems[T any](page, pageSize int) ([]T, error) {
	var entities []T

	db := initializers.DB

	switch any(new(T)).(type) {
	case *models.Menu:
		db = db.Preload("Products")
	}

	err := db.
		Order("id ASC").
		Limit(pageSize).
		Offset(page).
		Find(&entities).
		Error

	if err != nil {
		return entities, err
	}

	return entities, nil
}

func DeleteSellableItemByID[T any](id int) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	return initializers.DB.Delete(new(T), id).Error
}

func CreateSellableItem[T any](entity *T) (*T, error) {
	if err := initializers.DB.Create(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}
