package Queries

import (
	"awesomeKonstru/backend/models"

	"gorm.io/gorm"
)

func QueryActivityByID(db *gorm.DB, id string) (models.Actividad, error) {
	item := models.Actividad{}
	if err := db.First(&item, "id = ?", id).Error; err != nil {
		return models.Actividad{}, err
	}
	return item, nil
}
func QueryActividadByNombre(db *gorm.DB, nombre string) ([]models.Actividad, error) {
	var items []models.Actividad
	if err := db.Where("descripcion LIKE ?", "%"+nombre+"%").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
func QueryActivityByUserID(db *gorm.DB, id string) ([]models.Actividad_Usuario, error) {
	var items []models.Actividad_Usuario
	if err := db.Where("usuario_id = ?", id).Find(&items).Error; err != nil {
		return []models.Actividad_Usuario{}, err
	}
	return items, nil
}

func QueryActividadUserByNombre(db *gorm.DB, nombre string, id string) ([]models.Actividad_Usuario, error) {
	var items []models.Actividad_Usuario
	if err := db.Where("descripcion LIKE ?", "%"+nombre+"%").Where("usuario_id = ?", id).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func QueryActividadUserById(db *gorm.DB, idu string, id string) (models.Actividad_Usuario, error) {
	var item models.Actividad_Usuario
	if err := db.Where("id_actividad_usuario = ? AND usuario_id = ?", id, idu).First(&item).Error; err != nil {
		return models.Actividad_Usuario{}, err
	}
	return item, nil
}

func InsertInToActivitiesUser(db *gorm.DB, activityUser models.Actividad_Usuario) (error, int) {
	var lastInsertedID int
	if err := db.Create(&activityUser).Error; err != nil {
		return err, -1
	}
	var activityUser2 models.Actividad_Usuario
	if err := db.Last(&activityUser2).Error; err != nil {
		// Manejar el error aquí
		return err, -1
	} else {
		lastInsertedID = activityUser2.ID_Actividad_Usuario
		// Utilizar lastInsertedID
	}
	return nil, lastInsertedID
}
func InsertInToActivities(db *gorm.DB, activity models.Actividad) error {
	if err := db.Create(&activity).Error; err != nil {
		return err
	}
	return nil
}

func DeleteActivityByID(db *gorm.DB, id string) error {
	if err := db.Where("id = ?", id).Delete(&models.Actividad{}).Error; err != nil {
		return err
	}
	return nil
}
