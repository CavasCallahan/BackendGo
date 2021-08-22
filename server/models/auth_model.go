package models

type AuthModel struct {
	Base
	Email    string `json:"email" gorm:"type:varchar(255);unique_index"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	IsValid  bool   `json:"is_valid"`
}

// func (auth_model *AuthModel) BeforeCreate(rhx *gorm.DB) (err error) {
// 	auth_model.Password = services.SHA256Encoder(auth_model.Password) //encript's the user password everytime that creates a new user
// 	return
// }
