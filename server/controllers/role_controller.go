package controller

import (
	"fmt"

	"github.com/CavasCallahan/firstGo/server/database"
	"github.com/CavasCallahan/firstGo/server/models"
	"github.com/gin-gonic/gin"
)

func CreateRole(context *gin.Context) {

	db := database.GetDataBase()

	var role *models.RolesModel

	if err := context.ShouldBindJSON(&role); err != nil {
		context.JSON(400, "Please provid valid json")
		return
	}

	dbCreateError := db.Create(&role).Error

	if dbCreateError != nil {
		context.JSON(500, dbCreateError)
		return
	}

	context.JSON(201, "Role created with success!")

}

func GetRoles(context *gin.Context) {

	var role *[]models.RolesModel
	db := database.GetDataBase()

	dbFindError := db.Find(&role).Error

	if dbFindError != nil {
		context.JSON(500, dbFindError)
		return
	}

	context.JSON(200, role)

}

func UpdateRole(context *gin.Context) {

	type UpdateRoleInfo struct {
		OldName string `json:"old_name"`
		NewName string `json:"new_name"`
	}

	db := database.GetDataBase()
	var role_info *UpdateRoleInfo

	if err := context.ShouldBindJSON(&role_info); err != nil {
		context.JSON(400, "Please provid valid json")
		return
	}

	role := models.RolesModel{
		RoleName: role_info.NewName,
	}
	dbFindError := db.Where("role_name = ?", role_info.OldName).Updates(&role).Error

	if dbFindError != nil {
		context.JSON(500, dbFindError)
		return
	}

	context.JSON(204, "Role updated with success!")
}

func DeleteRole(context *gin.Context) {

	db := database.GetDataBase()
	var role_info *models.RolesModel

	if err := context.ShouldBindJSON(&role_info); err != nil {
		context.JSON(400, "Please provid valid json")
		return
	}

	fmt.Print(role_info.RoleName)

	var role *models.RolesModel
	dbDeleteError := db.Where("role_name = ?", role_info.RoleName).Delete(&role).Error

	if dbDeleteError != nil {
		context.JSON(500, dbDeleteError)
		return
	}

	context.JSON(202, "The Role deleted with success!")

}

func SignRole(context *gin.Context) {

	type SignRoleInfo struct {
		Email    string `json:"email"`
		RoleName string `json:"role_name"`
	}

	db := database.GetDataBase()
	var sing_info *SignRoleInfo

	if err := context.ShouldBindJSON(&sing_info); err != nil {
		context.JSON(400, "Please provid valid json")
		return
	}

	//Find the auth by email
	var auth *models.AuthModel
	dbFindAuthError := db.Where("email = ?", sing_info.Email).First(&auth).Error

	if dbFindAuthError != nil {
		context.JSON(500, dbFindAuthError)
		return
	}

	//Find the role by name
	var role *models.RolesModel
	dbFindRoleError := db.Where("role_name = ?", sing_info.RoleName).First(&role).Error

	if dbFindRoleError != nil {
		context.JSON(500, dbFindRoleError)
		return
	}

	role_user_model := models.RoleUserModel{
		AuthId: auth.ID,
		RoleId: role.ID,
	}

	dbCreateRoleUserError := db.Create(&role_user_model).Error

	if dbCreateRoleUserError != nil {
		context.JSON(500, dbCreateRoleUserError)
		return
	}

	context.JSON(200, "The user "+sing_info.Email+" was sign to the role "+sing_info.RoleName)
}

func SingOutRole(context *gin.Context) {

	db := database.GetDataBase()

	type SignOutInfo struct {
		Email    string `json:"email"`
		RoleName string `json:"role_name"`
	}

	var singout_info *SignOutInfo

	if err := context.ShouldBindJSON(&singout_info); err != nil {
		context.JSON(400, "Please provid valid json")
		return
	}

	//find the auth model
	var auth_model *models.AuthModel
	dbFindAuthError := db.Where("email = ?", singout_info.Email).First(&auth_model).Error

	if dbFindAuthError != nil {
		context.JSON(500, dbFindAuthError)
		return
	}

	var role *models.RolesModel
	dbFindRoleError := db.Where("role_name = ?", singout_info.RoleName).First(&role).Error

	if dbFindRoleError != nil {
		context.JSON(500, dbFindRoleError)
		return
	}

	//find role user model
	var role_user *models.RoleUserModel
	dbFindRoleUserError := db.Where("auth_id = ?", auth_model.ID).First(&role_user).Error

	if dbFindRoleError != nil {
		context.JSON(500, dbFindRoleUserError)
		return
	}

	role_user.RoleId = role.ID //Changes the role of the user

	dbUpdateError := db.Updates(&role_user).Error

	if dbUpdateError != nil {
		context.JSON(500, dbUpdateError)
		return
	}

	context.JSON(204, "Role Updated!")
}
