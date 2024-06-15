package Queries

import (
	"awesomeKonstru/backend/models"
	"gorm.io/gorm"
)

func InsertNewProject(db *gorm.DB, project *models.Proyectos) (*models.Proyectos, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	if err := tx.Create(project).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return project, nil
}

func GetProjectsByUserID(db *gorm.DB, userID string) ([]models.Proyectos, error) {
	var proyectos []models.Proyectos
	if err := db.Where("usuario_id = ?", userID).Order("created_at desc").Find(&proyectos).Error; err != nil {
		return nil, err
	}
	return proyectos, nil
}

func GetProjectByID(db *gorm.DB, ID string) (*models.Proyectos, error) {
	var proyecto models.Proyectos
	if err := db.Where("id_proyecto = ?", ID).First(&proyecto).Error; err != nil {
		return nil, err
	}
	return &proyecto, nil
}

func InsertNewProjectActivities(db *gorm.DB, projectActivities []models.Proyectos_actividades) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Create(&projectActivities).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
