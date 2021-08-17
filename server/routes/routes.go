package routes

import (
	"github.com/CavasCallahan/firstGo/server/controller"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("api/")
	{
		auth := main.Group("auth")
		{
			auth.GET("/", controller.Login)
		}
	}

	return router
}
