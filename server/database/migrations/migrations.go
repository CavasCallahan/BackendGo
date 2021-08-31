package migrations

import (
	"github.com/CavasCallahan/firstGo/server/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.Migrator().DropTable(models.AuthModel{})
	db.Migrator().DropTable(models.UserModel{})
	db.Migrator().DropTable(models.TokenModel{})
	db.Migrator().DropTable(models.RolesModel{})
	db.Migrator().DropTable(models.RoleUserModel{})
	db.Migrator().DropTable(models.NewsModel{})
	db.AutoMigrate(models.AuthModel{})
	db.AutoMigrate(models.UserModel{})
	db.AutoMigrate(models.TokenModel{})
	db.AutoMigrate(models.RolesModel{})
	db.AutoMigrate(models.RoleUserModel{})
	db.AutoMigrate(models.NewsModel{})
}
