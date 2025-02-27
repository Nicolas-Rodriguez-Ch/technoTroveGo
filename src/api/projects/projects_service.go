package projects

import (
	// "fmt"
	"technoTroveServer/src/models"

	"gorm.io/gorm"
)

func getAllProjects(db *gorm.DB) ([]models.ProjectResponse, error) {
	var projects []models.ProjectResponse
	err := db.Model(&models.Project{}).
		Select("projects.id, projects.title, projects.description, projects.images, projects.links, projects.user_id, users.full_name").
		Joins("JOIN users ON users.id = projects.user_id").
		Where("projects.active = ?", true).
		Find(&projects).Error

	return projects, err
}

func getProjectById(id string, db *gorm.DB) (*models.ProjectResponse, error) {
	var project models.ProjectResponse
	err := db.Model(&models.Project{}).
		Select("projects.id, projects.title, projects.description, projects.images, projects.links, projects.user_id, users.full_name, users.email, users.description, users.contact_info, users.profile_picture").
		Joins("JOIN users ON users.id = projects.user_id").
		Where("projects.id = ? AND projects.active = ?", id, true).
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
