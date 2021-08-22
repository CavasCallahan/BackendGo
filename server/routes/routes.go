package routes

import (
	controller "github.com/CavasCallahan/firstGo/server/controllers"
	"github.com/CavasCallahan/firstGo/server/middlewares"
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
		role := main.Group("role")
		{
			role.POST("/", middlewares.AuthMiddleware([]string{"manager", "user"}), controller.CreateRole)       //Handles the creation of the role
			role.GET("/", middlewares.AuthMiddleware([]string{"admin", "manager", "user"}), controller.GetRoles) //Handles selection of the role
			role.PUT("/", middlewares.AuthMiddleware([]string{"admin", "manager"}), controller.UpdateRole)       //Handles the update of the role
			role.DELETE("/", middlewares.AuthMiddleware([]string{"admin", "manager"}), controller.DeleteRole)    //Handles the delete of the role
			role.POST("/signrole", middlewares.AuthMiddleware([]string{"admin", "manager", "user"}), controller.SignRole)
		}
	}

	return router
}
