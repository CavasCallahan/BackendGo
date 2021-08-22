package models

type RoleUserModel struct {
	Base
	RoleId string `json:"role_id"`
	AuthId string `json:"auth_id"`
}
