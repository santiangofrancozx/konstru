package Query

import (
	"awesomeKonstru/backend/config"
	"awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
	"gorm.io/gorm"
)

func QueryUserByUsername(db *gorm.DB, email string) (models.Usuario, error) {
	user := models.Usuario{}
	if err := db.First(&user, "email = ?", email).Error; err != nil {
		return models.Usuario{}, err
	}
	return user, nil
}

func SelectUserByUsername(email string) (models.Usuario, error) {
	dsn := config.DB_DSN
	db, err := Connection_Migrates.Connect(dsn)
	if err != nil {
		return models.Usuario{}, err
	}
	defer Connection_Migrates.Disconnect(db)

	return QueryUserByUsername(db, email)
}
