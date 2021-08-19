package controller

import (
	"github.com/CavasCallahan/firstGo/server/database"
	"github.com/CavasCallahan/firstGo/server/models"
	"github.com/CavasCallahan/firstGo/server/services"
	"github.com/gin-gonic/gin"
)

func CreateInformation(context *gin.Context) {
	var user *models.UserModel
	db := database.GetDataBase()

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(401, "Invalid json provided")
		return
	}

	tokenAuth, err := services.ExtractTokenMetaData(context.Request)

	if err != nil {
		context.JSON(401, "Unauthorized")
		return
	}

	user.AuthId = tokenAuth.AuthId

	//Save the user information on the database
	db_err := db.Create(&user).Error

	if db_err != nil {
		context.JSON(500, db_err)
		return
	}

	context.JSON(200, user)
}

func GetInformation(context *gin.Context) {
	var user *models.UserModel
	db := database.GetDataBase()

	tokenAuth, err := services.ExtractTokenMetaData(context.Request)

	if err != nil {
		context.JSON(401, "Unauthorized")
		return
	}

	auth_id := tokenAuth.AuthId

	dbError := db.Where("auth_id = ?", auth_id).First(&user).Error

	if dbError != nil {
		context.JSON(400, dbError)
		return
	}

	context.JSON(200, user)
}

func UpdateInformation(context *gin.Context) {

	var update_user models.UserModel
	db := database.GetDataBase()

	if err := context.ShouldBindJSON(&update_user); err != nil {
		context.JSON(401, "Invalid json provided")
	}

	tokenAuth, err := services.ExtractTokenMetaData(context.Request)

	if err != nil {
		context.JSON(401, "Unauthorized")
		return
	}

	dbError := db.Where("auth_id = ?", tokenAuth.AuthId).Updates(&update_user).Error

	if dbError != nil {
		context.JSON(500, dbError)
		return
	}

	context.JSON(200, "User updated")
}
