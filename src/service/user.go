package service

import (
	"errors"
	"wackdo/src/initializers"
	"wackdo/src/models"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := initializers.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, &EntityNotFoundError{models.User{}}
		}
		return user, err
	}
	return user, nil
}

func GetUserById(id uint) (models.User, error) {
	var user models.User
	err := initializers.DB.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, &EntityNotFoundError{models.User{}}
		}
		return user, err
	}
	return user, nil
}

func GetUsers(page, pageSize int) ([]models.User, error) {
	var users []models.User

	err := initializers.DB.
		Limit(pageSize).
		Offset(page * pageSize).
		Find(&users).
		Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func CreateUser(user models.User) (models.User, error) {
	if err := initializers.DB.Create(&user).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.User{}, &DuplicateEmailError{}
		}
		return models.User{}, err
	}
	return user, nil
}

func UpdateUserFull(user models.User) (models.User, error) {
	err := initializers.DB.Save(&user).Error
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return user, &DuplicateEmailError{}
		}
		return user, err
	}
	return user, nil
}
