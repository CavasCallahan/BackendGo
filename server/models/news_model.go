package models

type NewsModel struct {
	Base
	Title       string `json:"title" gorm:"type:varchar(100)"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
	Author      string `json:"author" gorm:"type:varchar(50)"`
}
