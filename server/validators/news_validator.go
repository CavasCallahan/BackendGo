package validators

import (
	"net/url"

	"github.com/CavasCallahan/firstGo/server/models"
)

func NewsValidator(news *models.NewsModel) string {

	if len(news.Title) < 1 {
		return "You must provid a news name"
	}

	if len(news.Description) < 1 {
		return "You must provid a description"
	}

	if len(news.Author) < 1 {
		return "You must provid a author"
	}

	if !(len(news.Author) > 5 && len(news.Author) < 16) {
		return "The Author name must be between 5 and 16 characters"
	}

	_, err := url.ParseRequestURI(news.Thumbnail)

	if err != nil {
		return "Please provid a correct thumbnail"
	}

	return ""
}
