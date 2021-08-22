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

func RoleUserSeeder() {

	var role *models.RolesModel
	dbFindRoleError := db.Where("role_name = ?", "admin").First(&role).Error

	if dbFindRoleError != nil {
		return
	}

	var auth *models.AuthModel
	dbFindAuthError := db.Where("email = ?", "fire@deadshot.king").First(&auth).Error

	if dbFindAuthError != nil {
		return
	}

	role_user_model := models.RoleUserModel{
		AuthId: auth.ID,
		RoleId: role.ID,
	}

	dbCreateError := db.Create(&role_user_model).Error

	if dbCreateError != nil {
		return
	}
}

func PopulateDatabase() {
	AuthSeeder()
	RoleSeeder()
	RoleUserSeeder()
}
