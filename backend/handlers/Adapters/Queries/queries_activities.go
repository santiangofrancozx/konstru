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
func QueryActivityByOwnerID(db *gorm.DB, id string) ([]models.Actividad, error) {
	var items []models.Actividad
	if err := db.Where("created_by = ?", id).Find(&items).Error; err != nil {
		return []models.Actividad{}, err
	}
	return items, nil
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
func InsertInToActivities(db *gorm.DB, activity models.Actividad) (error, string) {
	// Intentar crear la actividad en la base de datos
	if err := db.Create(&activity).Error; err != nil {
		return err, "" // Devolver el error y una cadena vacía como ID
	}

	// El ID de la actividad debería estar actualizado en el modelo 'activity' después de la inserción
	return nil, activity.ID
}

func DeleteActivityByID(db *gorm.DB, id string, creator string) error {
	if err := db.Where("id = ?", id).Where("created_by = ?", creator).Delete(&models.Actividad{}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateActividad(db *gorm.DB, actividad models.Actividad) error {
	// Buscar la actividad por su ID
	var existingActividad models.Actividad
	if err := db.Where("id = ?", actividad.ID).First(&existingActividad).Error; err != nil {
		return err // Manejar el error si no se encuentra la actividad
	}

	// Actualizar los campos necesarios
	existingActividad.Descripcion = actividad.Descripcion
	existingActividad.Unidad = actividad.Unidad
	existingActividad.PrecioBase = actividad.PrecioBase

	// Guardar los cambios en la base de datos
	if err := db.Save(&existingActividad).Error; err != nil {
		return err // Manejar el error si falla el guardado
	}

	return nil
}
