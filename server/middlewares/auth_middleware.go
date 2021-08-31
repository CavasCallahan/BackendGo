package middlewares

import (
	"github.com/CavasCallahan/firstGo/server/database"
	"github.com/CavasCallahan/firstGo/server/models"
	"github.com/CavasCallahan/firstGo/server/services"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(role_names []string) gin.HandlerFunc {
	return func(context *gin.Context) {

		db := database.GetDataBase()

		tokenAuth, err := services.ExtractTokenMetaData(context.Request)

		if err != nil {
			context.JSON(401, "Please Provid a valid Token")
			context.AbortWithStatus(401)
		}
		//Find's the role_id
		var role_user_model *models.RoleUserModel
		dbFindRoleUserError := db.Where("auth_id", tokenAuth.AuthId).First(&role_user_model).Error

		if dbFindRoleUserError != nil {
			context.AbortWithStatus(500)
		}

		//Finds the role
		var role *models.RolesModel

		dbFindRoleError := db.Where("id", role_user_model.RoleId).First(&role).Error

		if dbFindRoleError != nil {
			context.AbortWithStatus(500)
		}

		for i := 0; i < len(role_names); i++ {
			if role_names[i] == role.RoleName {
				context.Next()
				return
			}
		}

		context.AbortWithStatus(401)
	}
}
