package controller

import (
	"net/http"

	"github.com/CavasCallahan/firstGo/server/models"
	"github.com/gin-gonic/gin"
)

var auth_ = models.AuthModel{
	Email:    "fire@deadshot.com",
	Password: "123",
}

func Login(context *gin.Context) {
	var user_auth models.AuthModel

	if err := context.ShouldBindJSON(&user_auth); err != nil {
		context.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	if auth_.Email != user_auth.Email || auth_.Password != user_auth.Password {
		context.JSON(http.StatusUnauthorized, "Please provid a valid login")
		return
	}

	context.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
