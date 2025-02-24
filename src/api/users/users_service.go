package users

import (
	"errors"
	"fmt"
	"technoTroveServer/src/models"

	"gorm.io/gorm"
)

func getAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	err := db.Preload("Projects", "active = ?", true).Find(&users).Error
	return users, err
}

func CreateUser(input *models.User, db *gorm.DB) (*models.User, error) {
	fmt.Println("Esto es input:", input)
	result := db.Create(input)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, errors.New("a user with this email already exists")
		}
		return nil, result.Error
	}
	return input, nil
}
