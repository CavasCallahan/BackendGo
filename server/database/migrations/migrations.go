package migrations

import (
	"github.com/CavasCallahan/firstGo/server/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.AuthModel{})
	db.AutoMigrate(models.UserModel{})
}
