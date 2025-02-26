package users

import (
	"errors"
	"technoTroveServer/src/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func getAllUsers(db *gorm.DB) ([]models.UserResponse, error) {
	var users []models.UserResponse
	err := db.Model(&models.User{}).
		Select("id, full_name, email, description, contact_info, profile_picture").
		Preload("Projects", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, active, title, description, images, links").Where("active = ?", true)
		}).
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
	err := db.Preload("Projects", "active = ?", true).
		Where("id = ?", id).
		First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func getUserProfile(id string, db *gorm.DB) (*models.UserResponse, error) {
	var user models.UserResponse
	err := db.Model(&models.User{}).
		Select("id, full_name, email, description, contact_info, profile_picture").
		Preload("Projects", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, active, title, description, images, links").Where("active = ?", true)
		}).
		Where("id = ?", id).
		First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func updateUser(id string, input *models.User, db *gorm.DB) (*models.UserResponse, error) {
	var existingUser models.User

	if err := db.First(&existingUser, "id = ?", id).Error; err != nil {
		return nil, errors.New("user not found")
	}
	updates := map[string]interface{}{}

	if input.FullName != "" {
		updates["full_name"] = input.FullName
	}
	if input.Email != "" {
		updates["email"] = input.Email
	}
	if input.Password != "" {
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, errors.New("error hashing password")
		}
		updates["password"] = string(hashedPass)
	}
	if input.Description != "" {
		updates["description"] = input.Description
	}
	if len(input.ContactInfo) > 0 {
		updates["contact_info"] = input.ContactInfo
	}

	if input.ProfilePicture != nil && (existingUser.ProfilePicture == nil || *input.ProfilePicture != *existingUser.ProfilePicture) {
		updates["profile_picture"] = input.ProfilePicture
	}

	if len(updates) == 0 {
		return nil, errors.New("no updates provided")
	}

	if err := db.Model(&existingUser).Updates(updates).Error; err != nil {
		return nil, err
	}

	var updatedUser models.UserResponse
	if err := db.Model(&models.User{}).
		Select("id, full_name, email, description, contact_info, profile_picture").
		Where("id = ?", id).
		First(&updatedUser).Error; err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func deactivateUser(id string, db *gorm.DB) (*models.UserResponse, error) {
	if err := db.Where("user_id = ?", id).Delete(&models.Project{}).Error; err != nil {
		return nil, err
	}

	var deletedUser models.UserResponse
	if err := db.Model(&models.User{}).
		Select("id, full_name, email, description, contact_info, profile_picture").
		Where("id = ?", id).
		First(&deletedUser).Error; err != nil {
		return nil, err
	}

	if err := db.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return nil, err
	}
	return &deletedUser, nil
}
