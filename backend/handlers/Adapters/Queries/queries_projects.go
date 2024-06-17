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
	if err := db.Preload("Usuario").Where("usuario_id = ?", userID).Order("created_at desc").Find(&proyectos).Error; err != nil {
		return nil, err
	}
	return proyectos, nil
}

func GetProjectByID(db *gorm.DB, ID string) (*models.Proyectos, error) {
	var proyecto models.Proyectos
	if err := db.Preload("Usuario").Where("id_proyecto = ?", ID).First(&proyecto).Error; err != nil {
		return nil, err
	}
	return &proyecto, nil
}

func InsertNewProjectActivities(db *gorm.DB, projectActivities []models.Proyectos_actividades) error {
	if err := db.Create(projectActivities).Error; err != nil {
		return err
	}
	return nil
}

func GetAllProjectActivities(db *gorm.DB) ([]models.Proyectos_actividades, error) {
	var items []models.Proyectos_actividades
	if err := db.Preload("Proyectos").Preload("Actividad").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
