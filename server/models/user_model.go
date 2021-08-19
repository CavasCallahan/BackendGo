package models

type UserModel struct {
	Base
	AuthId   string `json:"auth_id"`
	UserName string `json:"username"`
	LastName string `json:"lastname"`
	Avatar   string `json:"avatar"`
	IsMember bool   `json:"is_member"`
}
