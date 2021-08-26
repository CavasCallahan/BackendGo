package validators

import (
	"net/url"

	"github.com/CavasCallahan/firstGo/server/models"
)

func UserValidator(user_model *models.UserModel) string {

	if !(len(user_model.UserName) > 8) {
		return "The user_name must have at least 8 characters"
	}

	if !(len(user_model.LastName) > 5) {
		return "The last_name must have at least 8 characters"
	}

	_, err := url.ParseRequestURI(user_model.Avatar)

	if err != nil {
		return "Please provid a correct avatar"
	}

	return ""

}
