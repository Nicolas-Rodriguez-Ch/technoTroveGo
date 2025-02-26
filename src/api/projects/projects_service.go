package projects

import (
	"technoTroveServer/src/models"

	"gorm.io/gorm"
)

func getAllProjects(db *gorm.DB) ([]models.ProjectResponse, error) {
	var projects []models.ProjectResponse
	err := db.Model(&models.Project{}).
		Select("id, title, description, images, links, user_id").
		Where("active = ?", true).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, full_name, email, contact_info, profile_picture")
		}).
		Find(&projects).Error

	return projects, err
}
