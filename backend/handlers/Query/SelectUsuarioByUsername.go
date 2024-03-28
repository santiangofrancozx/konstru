package Query

import (
	"awesomeKonstru/backend/handlers"
	"awesomeKonstru/backend/models"
	"gorm.io/gorm"
)

func QueryUserByUsername(db *gorm.DB, username string) (models.Usuario, error) {
	item := models.Usuario{}
	if err := db.First(&item, username).Error; err != nil {
		return models.Usuario{}, err
	}
	return item, nil
}

func SelectUserByUsername(dsn string, username string) (models.Usuario, error) {
	db, err := handlers.Connect(dsn)
	if err != nil {
		return models.Usuario{}, err
	}
	defer handlers.Disconnect(db)

	return QueryUserByUsername(db, username)
}
