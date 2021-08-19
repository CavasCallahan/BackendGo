package models

type AuthModel struct {
	Base
	Email    string `json:"email" gorm:"type:varchar(255);unique_index"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	IsValid  bool   `json:"is_valid"`
}
