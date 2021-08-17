package models

type AuthModel struct {
	Base
	Email        string
	Password     string
	RefreshToken string
}
