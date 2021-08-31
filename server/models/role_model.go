package models

type RolesModel struct {
	Base
	RoleName string `json:"role_name" gorm:"type:varchar(10);unique"`
}
