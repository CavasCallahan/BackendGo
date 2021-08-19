package controller

import (
	"github.com/CavasCallahan/firstGo/server/database"
	"github.com/CavasCallahan/firstGo/server/models"
	"github.com/CavasCallahan/firstGo/server/services"
	"github.com/gin-gonic/gin"
)

func CreateInformation(context *gin.Context) {
	var user models.UserModel
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
