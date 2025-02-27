package projects

import (
	// "fmt"
	"errors"
	"reflect"
	"technoTroveServer/src/models"

	"gorm.io/gorm"
)

func getAllProjects(db *gorm.DB) ([]models.ProjectResponse, error) {
	var projects []models.ProjectResponse
	err := db.Model(&models.Project{}).
		Select("projects.id, projects.title, projects.description, projects.images, users.full_name").
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
func updateProject(id string, input *models.Project, db *gorm.DB) (*models.ProjectResponse, error) {
	var existingProject models.Project

	err := db.Where("user_id = ?", id).First(&existingProject).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("project not found")
	}
	if err != nil {
		return nil, err
	}

	updateData := map[string]interface{}{}

	if input.Title != "" {
		updateData["title"] = input.Title
	}
	if input.Description != "" {
		updateData["description"] = input.Description
	}
	if len(input.Images) > 0 {
		newImages := append(existingProject.Images, input.Images...)
		if !reflect.DeepEqual(newImages, existingProject.Images) {
			updateData["images"] = newImages
		}
	}
	if len(input.Links) > 0 {
		updateData["links"] = input.Links
	}

	if len(updateData) > 0 {
		err = db.Model(&existingProject).Updates(updateData).Error
		if err != nil {
			return nil, err
		}
	}

	var updatedProject models.ProjectResponse
	err = db.Model(&models.Project{}).
		Select("projects.id, projects.title, projects.description, projects.images, projects.links, projects.user_id, users.full_name, users.email, users.description, users.contact_info, users.profile_picture").
		Joins("JOIN users ON users.id = projects.user_id").
		Where("projects.id = ?", existingProject.ID).
		First(&updatedProject).Error

	if err != nil {
		return nil, err
	}

	return &updatedProject, nil
}
