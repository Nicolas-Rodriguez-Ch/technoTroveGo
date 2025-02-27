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
			return db.Select("full_name")
		}).
		Find(&projects).Error

	return projects, err
}

func getProjectById(id string, db *gorm.DB) (*models.ProjectResponse, error) {
	var project models.ProjectResponse
	err := db.Model(&models.Project{}).
		Where("id = ? AND active = ?", id, true).
		Select("id, title, description, images, links").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, full_name, email, contact_info, profile_picture")
		}).
		First(&project).Error

	if err != nil {
		return nil, err
	}
	return &project, nil
}

func createProject(input *models.Project, db *gorm.DB) (*models.Project, error) {
	result := db.Create(input)
	if result.Error != nil {
		return nil, result.Error
	}
	return input, nil
}
