package validators

import (
	"regexp"

	"github.com/CavasCallahan/firstGo/server/database"
	"github.com/CavasCallahan/firstGo/server/models"
)

func VerifyEmail(email string) bool {
	emailRegexp := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return !emailRegexp.MatchString(email)
}

func AuthValidation(auth *models.AuthModel) string {
	db := database.GetDataBase()

	passNumberRegexp := regexp.MustCompile("[0-9]")
	passUpperRegexp := regexp.MustCompile("[A-Z]")

	if len(auth.Email) < 1 {
		return "Please provid email"
	}

	if VerifyEmail(auth.Email) {
		return "Please provid a valid email"
	}

	dbFindError := db.Where("email", auth.Email).Error

	if dbFindError != nil { // Fix's this please
		return "The email is in use"
	}

	if !(len(auth.Password) > 8 && len(auth.Password) < 16) {
		return "The passoword have to have at least 8 to 16 words or numbers"
	}

	if !passNumberRegexp.MatchString(auth.Password) {
		return "The password must have one number number"
	}

	if !passUpperRegexp.MatchString(auth.Password) {
		return "The password must have one letter uppercase"
	}

	return ""
}
