package users

import (
	"technoTroveServer/src/models"

	"gorm.io/gorm"
)

func getAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	err := db.Preload("Projects", "active = ?", true).Find(&users).Error
	return users, err
}
