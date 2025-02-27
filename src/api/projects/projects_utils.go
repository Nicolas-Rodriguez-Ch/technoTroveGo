package projects

import (
	"errors"
	"net/http"
	"technoTroveServer/src/models"

	"gorm.io/gorm"
)

func checkProjectOwnerShip(userId string, projectId string, db *gorm.DB) (int, error) {
	var project models.Project

	err := db.Select("user_id").
		Where("id = ?", projectId).
		First(&project).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound, err
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if project.UserID != userId {
		return http.StatusForbidden, err
	}
	return http.StatusOK, nil
}
