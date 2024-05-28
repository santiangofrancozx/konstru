package Queries

import (
	"awesomeKonstru/backend/models"
	"gorm.io/gorm"
)

func InsertNewProject(db *gorm.DB, project *models.Proyectos) error {
	if err := db.Create(project).Error; err != nil {
		return err
	}
	return nil
}

func GetProjectsByUserID(db *gorm.DB, userID string) ([]models.Proyectos, error) {
	var proyectos []models.Proyectos
	if err := db.Where("usuario_id = ?", userID).Order("created_at desc").Find(&proyectos).Error; err != nil {
		return nil, err
	}
	return proyectos, nil
}
