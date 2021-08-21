package controller

import (
	"net/http"
	"time"

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

	context.JSON(200, "The user was created!")
}

func ForgotPassword(context *gin.Context) {

	type ResetInfo struct {
		Email string `json:"email"`
	}

	var reset_info ResetInfo

	db := database.GetDataBase()

	if err := context.ShouldBindJSON(&reset_info); err != nil {
		context.JSON(400, "Please provid a valid information")
		return
	}

	var auth *models.AuthModel
	dbFindError := db.Where("email = ?", reset_info.Email).First(&auth).Error

	if dbFindError != nil {
		context.JSON(500, dbFindError)
		return
	}

	token := services.GenerateStaticToken()

	token_model := models.TokenModel{
		AuthId: auth.ID,
		Value:  services.SHA256Encoder(token),
		Type:   "reset_password",
	}

	dbCreateError := db.Create(&token_model).Error

	if dbCreateError != nil {
		context.JSON(500, dbCreateError)
		return
	}

	//Send Email with the instruction

	context.JSON(200, token)
}

type ValidateCredentials struct {
	NewPassword      string `json:"new_password"`
	ComfirmPassoword string `json:"confirm_password"`
}

func ValidateResetEmail(context *gin.Context) {
	var auth *ValidateCredentials
	db := database.GetDataBase()

	if err := context.ShouldBindJSON(&auth); err != nil { // get's the json information
		context.JSON(400, "Please provid a valid information")
		return
	}

	if auth.NewPassword != auth.ComfirmPassoword {
		context.JSON(400, "The password and confirm password must be equal!") //confirms that password and confirm password are equal
		return
	}

	token, ok := context.GetQuery("token") //get's the token from query

	if !ok {
		context.JSON(400, "Token not Provided")
		return
	}

	var token_model *models.TokenModel
	dbFindTokenError := db.Where("value = ?", services.SHA256Encoder(token)).First(&token_model).Error //seartch the token in the database

	if dbFindTokenError != nil {
		context.JSON(500, dbFindTokenError)
		return
	}

	if token_model.Type != "reset_password" { //verify what type is the token
		context.JSON(404, "Token is not valid")
		return
	}

	row_date := time.Now().Minute() - token_model.CreatedAt.Minute() //calculates how mutch time has passed

	if row_date > 15 { //15 minutes
		context.JSON(400, "Token expired")
		db.Delete(&token_model) //deletes the token
		return
	}

	var user_auth *models.AuthModel

	dbFindAuthError := db.Where("id = ?", token_model.AuthId).First(&user_auth).Error //find's who belong the token

	if dbFindAuthError != nil {
		context.JSON(500, dbFindTokenError)
		return
	}

	user_auth.Password = services.SHA256Encoder(auth.NewPassword) //changes and encodes the new password

	dbUpdateAuthError := db.Where("id = ?", user_auth.ID).Updates(&user_auth).Error //updates the user info

	if dbUpdateAuthError != nil {
		context.JSON(500, dbUpdateAuthError)
		return
	}

	dbDeleteError := db.Delete(&token_model).Error //deletes the token

	if dbDeleteError != nil {
		context.JSON(500, dbDeleteError)
		return
	}

	context.JSON(200, "Password changed!")

}
