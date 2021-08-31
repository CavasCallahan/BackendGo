package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/CavasCallahan/firstGo/server/database"
	"github.com/CavasCallahan/firstGo/server/models"
	"github.com/CavasCallahan/firstGo/server/services"
	"github.com/CavasCallahan/firstGo/server/validators"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

	context.JSON(201, gin.H{
		"acess_token":   token.AcessToken,
		"refresh_token": token.RefreshToken,
	})
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

	err := validators.AuthValidation(&models.AuthModel{
		Email:    sing_up_info.Email,
		Password: sing_up_info.ConfirmPassword,
	})

	if len(err) > 0 {
		context.JSON(400, err)
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
			"error": db_err.Error(),
		})
		return
	}

	var user_role *models.RolesModel
	dbFindUserRole := db.Where("role_name = ?", "user").First(&user_role).Error

	if dbFindUserRole != nil {
		context.JSON(500, gin.H{
			"error": db_err.Error(),
		})
		return
	}

	role_user := models.RoleUserModel{
		RoleId: user_role.ID,
		AuthId: auth.ID,
	}

	dbCreateUserRole := db.Create(&role_user).Error

	if dbCreateUserRole != nil {
		context.JSON(500, gin.H{
			"error": db_err.Error(),
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

	not_ok := validators.VerifyEmail(reset_info.Email)

	if not_ok {
		context.JSON(400, "Please provid a valid email")
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

func EmailVerification(context *gin.Context) {

	db := database.GetDataBase()

	type EmailVerificationInfo struct {
		Email string `json:"email"`
	}

	var email_verification_info *EmailVerificationInfo

	if err := context.ShouldBindJSON(&email_verification_info); err != nil {
		context.JSON(400, "Please provid valid json")
		return
	}

	not_ok := validators.VerifyEmail(email_verification_info.Email)

	if not_ok {
		context.JSON(400, "Please provid a valid email")
		return
	}

	var auth *models.AuthModel
	dbFindAuthError := db.Where("email = ?", email_verification_info.Email).First(&auth).Error

	if dbFindAuthError != nil {
		context.JSON(500, dbFindAuthError)
		return
	}

	token := services.GenerateStaticToken()

	token_model := models.TokenModel{
		AuthId: auth.ID,
		Value:  services.SHA256Encoder(token),
		Type:   "validation_email",
	}

	dbCreateError := db.Create(&token_model).Error

	if dbCreateError != nil {
		context.JSON(500, dbCreateError)
		return
	}

	//send Email

	context.JSON(202, token)
}

type ValidateCredentials struct {
	NewPassword      string `json:"new_password"`
	ComfirmPassoword string `json:"confirm_password"`
}

func ValidateResetPassword(context *gin.Context) {
	var auth *ValidateCredentials
	db := database.GetDataBase()

	if err := context.ShouldBindJSON(&auth); err != nil { // get's the json information
		context.JSON(400, "Please provid a valid information")
		return
	}

	err := validators.ValidatePassword(models.AuthModel{
		Password: auth.NewPassword,
	})

	if len(err) > 0 {
		context.JSON(400, err)
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

func RefreshToken(context *gin.Context) {
	mapToken := map[string]string{}

	if err := context.ShouldBindJSON(&mapToken); err != nil {
		context.JSON(400, err.Error())
		return
	}

	refreshToken := mapToken["refresh_token"]

	token, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte("OnlyAKingCanKilAkingAndOnlyAkingCanBeKilledByAking"), nil
	})

	if err != nil {
		context.JSON(400, err)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {

		auth_id, ok := claims["auth_id"].(string)

		if !ok {
			context.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		token, createErr := services.GenerateToken(auth_id)

		if createErr != nil {
			context.JSON(401, createErr.Error())
			return
		}

		context.JSON(201, gin.H{
			"acess_token":   token.AcessToken,
			"refresh_token": token.RefreshToken,
		})

	}
}

func ValidateEmail(context *gin.Context) {

	db := database.GetDataBase()

	token, ok := context.GetQuery("token") //get's the token from query

	if !ok {
		context.JSON(400, "Token not Provided")
		return
	}

	var token_model *models.TokenModel
	dbFindTokenError := db.Where("value = ?", services.SHA256Encoder(token)).First(&token_model).Error

	if dbFindTokenError != nil {
		context.JSON(500, dbFindTokenError.Error())
		return
	}

	if token_model.Type != "validation_email" {
		context.JSON(404, "Token is not valid")
		return
	}

	row_date := time.Now().Minute() - token_model.CreatedAt.Minute()

	if row_date > 15 {
		context.JSON(400, "Token expired")
		db.Delete(&token_model)
		return
	}

	var auth *models.AuthModel

	dbFindAuthError := db.Where("id = ?", token_model.AuthId).First(&auth).Error

	if dbFindAuthError != nil {
		context.JSON(500, dbFindAuthError.Error())
		return
	}

	auth.IsValid = true

	dbUpdateError := db.Updates(&auth).Error

	if dbUpdateError != nil {
		context.JSON(500, dbUpdateError.Error())
		return
	}

	dbDeleteError := db.Delete(&token_model).Error //deletes the token

	if dbDeleteError != nil {
		context.JSON(500, dbDeleteError)
		return
	}

	context.JSON(200, "Your account is now validated!")

}
