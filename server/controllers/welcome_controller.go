package controller

import "github.com/gin-gonic/gin"

func WelcomePage(context *gin.Context) {
	context.HTML(200, "welcome.html", nil)
}
