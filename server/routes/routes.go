package routes

import (
	controller "github.com/CavasCallahan/firstGo/server/controllers"
	"github.com/CavasCallahan/firstGo/server/middlewares"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("api/")
	{
		main.GET("/", controller.WelcomePage)
		router.NoRoute(controller.NotFoundRoute)
		auth := main.Group("auth")
		{
			auth.POST("/", controller.EmailVerification)
			auth.POST("/login", controller.Login)                           //Handles the login route
			auth.POST("/singup", controller.SingUp)                         //Handles the sing up route
			auth.POST("/forgot", controller.ForgotPassword)                 //Handles the forgot password route
			auth.POST("/forgot/validate", controller.ValidateResetPassword) //Handles the validation of forgot password route
			auth.GET("/validate", controller.ValidateEmail)                 //Handles the validation of the email
			auth.POST("/refresh_token", controller.RefreshToken)            //Generates a new acess_token
		}
		profile := main.Group("profile")
		{
			profile.POST("/", middlewares.AuthMiddleware([]string{"user", "manager", "admin"}), controller.CreateInformation)   //Handles the creation of the information of the user
			profile.GET("/", middlewares.AuthMiddleware([]string{"user", "manager", "admin"}), controller.GetInformation)       //Handles the selection of information
			profile.PUT("/", middlewares.AuthMiddleware([]string{"user", "manager", "admin"}), controller.UpdateInformation)    //Handles the update of the information
			profile.DELETE("/", middlewares.AuthMiddleware([]string{"user", "manager", "admin"}), controller.DeleteInformation) //Handles the delete of the information
		}
		role := main.Group("role")
		{
			role.POST("/", middlewares.AuthMiddleware([]string{"admin"}), controller.CreateRole)                          //Handles the creation of the role
			role.GET("/", middlewares.AuthMiddleware([]string{"admin", "manager", "user"}), controller.GetRoles)          //Handles selection of the role
			role.PUT("/", middlewares.AuthMiddleware([]string{"admin"}), controller.UpdateRole)                           //Handles the update of the role
			role.DELETE("/", middlewares.AuthMiddleware([]string{"admin"}), controller.DeleteRole)                        //Handles the delete of the role
			role.POST("/signrole", middlewares.AuthMiddleware([]string{"admin", "manager", "user"}), controller.SignRole) //Handles the sign to role
		}
		news := main.Group("news")
		{
			news.POST("/", middlewares.AuthMiddleware([]string{"admin", "manager"}), controller.CreateNews)
			news.GET("/", middlewares.AuthMiddleware([]string{"admin", "manager", "user"}), controller.GetNews)
			news.PUT("/", middlewares.AuthMiddleware([]string{"admin", "manager"}), controller.UpdateNews)
			news.DELETE("/", middlewares.AuthMiddleware([]string{"admin", "manager"}), controller.DeleteNews)
		}
	}

	return router
}
