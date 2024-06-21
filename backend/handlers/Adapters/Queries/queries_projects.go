package Queries

import (
	"awesomeKonstru/backend/models"
	"fmt"
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

func GetAllProjectActivities(db *gorm.DB, idp string) ([]models.Proyectos_actividades, error) {
	var items []models.Proyectos_actividades
	if err := db.Where("id_proyecto = ?", idp).Preload("Proyectos").Preload("Actividad").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func DeleteProjectOnCascade(db *gorm.DB, ID string) error {

	return nil
}

func DeleteProjectByID(db *gorm.DB, projectID string) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Where("id_proyecto = ?", projectID).Delete(&models.Proyectos{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func DeleteSpecificProjectActivities(db *gorm.DB, IDProyecto string, IDActivities []string) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer tx.Commit() // Asegura que la transacción se termine al finalizar la función

	// Eliminar las actividades específicas del proyecto
	result := tx.Where("id_proyecto = ?", IDProyecto).
		Where("id_actividad IN (?)", IDActivities).
		Delete(&models.Proyectos_actividades{})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	// Verificar si se eliminó al menos una fila
	if result.RowsAffected == 0 {
		tx.Rollback() // Rollback si no se eliminó ninguna fila
		return fmt.Errorf("no activities deleted for project %s and given IDs", IDProyecto)
	}

	return nil
}

func DeleteAllActivitiesProject(db *gorm.DB, IDProject string) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	result := tx.Where("id_proyecto = ?", IDProject).Delete(&models.Proyectos_actividades{})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if result.RowsAffected == 0 {
		tx.Rollback() // Rollback si no se eliminó ninguna fila
		return fmt.Errorf("no activities deleted for project %s and given IDs", IDProject)
	}

	return nil
}

func UpdateProjectActivities(db *gorm.DB, PAmodels []models.Proyectos_actividades, Idproject string) error {
	// Eliminar todas las actividades del proyecto para actualizar
	result := db.Where("id_proyecto = ?", Idproject).Delete(&models.Proyectos_actividades{})
	if result.Error != nil {
		return result.Error
	}
	// Insertar nuevas actividades si PAmodels no está vacío, si esta vacio es que se borro todas
	if len(PAmodels) > 0 {
		if err := db.Create(&PAmodels).Error; err != nil {
			return err
		}
	}
	return nil
}
