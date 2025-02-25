package users

import (
	"errors"
	"technoTroveServer/src/models"

	"gorm.io/gorm"
)

func getAllUsers(db *gorm.DB) ([]models.UserResponse, error) {
	var users []models.UserResponse
	err := db.Model(&models.User{}).
		Select("id, full_name, email, description, contact_info, profile_picture").
		Preload("Projects", "active = ?", true).
		Find(&users).Error
	return users, err
}

func CreateUser(input *models.User, db *gorm.DB) (*models.User, error) {
	result := db.Create(input)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, errors.New("a user with this email already exists")
		}
		return nil, result.Error
	}
	return input, nil
}

func getUserById(id string, db *gorm.DB) (*models.User, error) {
	var user models.User
	err := db.Preload("Projects", "active = ?", true).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
