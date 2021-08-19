package controller

import (
	"net/http"

	"github.com/CavasCallahan/firstGo/server/database"
	"github.com/CavasCallahan/firstGo/server/models"
	"github.com/CavasCallahan/firstGo/server/services"
	"github.com/gin-gonic/gin"
)

type SingUpUserInfo struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func Login(context *gin.Context) {
	db := database.GetDataBase()
	var user_auth models.AuthModel

	if err := context.ShouldBindJSON(&user_auth); err != nil {
		context.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	var user *models.AuthModel
	dbError := db.Where("email = ?", user_auth.Email).First(&user).Error

	if dbError != nil {
		context.JSON(500, dbError)
	}

	if user.Email != user_auth.Email || user.Password != services.SHA256Encoder(user_auth.Password) {
		context.JSON(http.StatusUnauthorized, "Please provid a valid login")
		return
	}

	token, err := services.GenerateToken(user.ID)

	if err != nil {
		context.JSON(500, err.Error()) //Gives the error
		return
	}

	context.JSON(201, token.AcessToken)
}

func SingUp(context *gin.Context) {
	db := database.GetDataBase()

	var sing_up_info SingUpUserInfo

	if err := context.ShouldBindJSON(&sing_up_info); err != nil {
		context.JSON(400, "Please provid a valid information")
		return
	}

	if sing_up_info.Password != sing_up_info.ConfirmPassword {
		context.JSON(400, "The password and confirm password has to be the same")
		return
	}

	sing_up_info.Password = services.SHA256Encoder(sing_up_info.Password)
	sing_up_info.ConfirmPassword = sing_up_info.Password

	//Create the user
	auth := models.AuthModel{Email: sing_up_info.Email, Password: sing_up_info.Password}

	//Save the user on the database
	db_err := db.Create(&auth).Error

	if db_err != nil {
		context.JSON(500, gin.H{
			"error": db_err,
		})
		return
	}

	context.JSON(200, "The user was created with sucess")
}
