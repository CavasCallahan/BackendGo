package database

import (
	"github.com/CavasCallahan/firstGo/server/models"
	"github.com/CavasCallahan/firstGo/server/services"
)

func AuthSeeder() {

	auth_model := models.AuthModel{
		Email:    "fire@deadshot.king",
		Password: services.SHA256Encoder("12345"),
	}

	dbErr := db.Create(&auth_model).Error

	if dbErr != nil {
		return
	}

}

func RoleSeeder() {

	auth_model := models.RolesModel{
		RoleName: "admin",
	}

	dbErr := db.Create(&auth_model).Error

	if dbErr != nil {
		return
	}

}

func PopulateDatabase() {
	AuthSeeder()
	RoleSeeder()
}
