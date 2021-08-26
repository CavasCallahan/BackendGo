package controller

import "github.com/gin-gonic/gin"

func NotFoundRoute(context *gin.Context) {
	context.HTML(404, "index.html", nil)
}
