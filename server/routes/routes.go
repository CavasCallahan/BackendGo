package routes

import (
	controller "github.com/CavasCallahan/firstGo/server/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("api/")
	{
		auth := main.Group("auth")
		{
			auth.POST("/login", controller.Login)           //Handles the login route
			auth.POST("/singup", controller.SingUp)         //Handles the sing up route
			auth.POST("/forgot", controller.ForgotPassword) //Handles the forgot password route
			auth.POST("/validate", controller.ValidateResetEmail)
		}
		profile := main.Group("profile")
		{
			profile.POST("/", controller.CreateInformation)   //Handles the creation of the information of the user
			profile.GET("/", controller.GetInformation)       //Handles the selection of information
			profile.PUT("/", controller.UpdateInformation)    //Handles the update of the information
			profile.DELETE("/", controller.DeleteInformation) //Handles the delete of the information
		}
	}

	return router
}
