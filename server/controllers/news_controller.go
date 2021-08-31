package controller

import (
	"github.com/CavasCallahan/firstGo/server/database"
	"github.com/CavasCallahan/firstGo/server/models"
	"github.com/CavasCallahan/firstGo/server/validators"
	"github.com/gin-gonic/gin"
)

func CreateNews(context *gin.Context) {

	db := database.GetDataBase()
	var news_model *models.NewsModel

	if err := context.ShouldBindJSON(&news_model); err != nil {
		context.JSON(400, "Please provid a valid json")
		return
	}

	err := validators.NewsValidator(news_model)

	if len(err) > 1 {
		context.JSON(404, err)
		return
	}

	dbCreateError := db.Create(&news_model).Error

	if dbCreateError != nil {
		context.JSON(500, dbCreateError.Error())
		return
	}

	context.JSON(201, "The news was created!")
}

func GetNews(context *gin.Context) {

	db := database.GetDataBase()

	var news []*models.NewsModel
	dbFindError := db.Find(&news).Error

	if dbFindError != nil {
		context.JSON(500, dbFindError)
		return
	}

	context.JSON(200, news)

}

type News_Info struct {
	Title       string `json:"title"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

func UpdateNews(context *gin.Context) {

	db := database.GetDataBase()

	var update_news *News_Info

	if err := context.ShouldBindJSON(&update_news); err != nil {
		context.JSON(400, "Please provid a valid json")
		return
	}

	id, ok := context.GetQuery("id") //get's the token from query

	if !ok {
		context.JSON(400, "id was not provided")
		return
	}

	new := models.NewsModel{
		Title:       update_news.Title,
		Thumbnail:   update_news.Thumbnail,
		Description: update_news.Description,
		Author:      update_news.Author,
	}

	err := validators.NewsValidator(&new)

	if len(err) > 1 {
		context.JSON(400, err)
		return
	}

	dbUpdateError := db.Model(&models.NewsModel{}).Where("id = ?", id).Updates(&new).Error

	if dbUpdateError != nil {
		context.JSON(500, dbUpdateError)
		return
	}

	context.JSON(204, "The News was been updated!")
}

func DeleteNews(context *gin.Context) {
	context.JSON(200, "Hello Delete")
}
