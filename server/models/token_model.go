package models

type TokenModel struct {
	Base
	AuthId string `json:"auth_id"`
	Value  string `json:"value"`
	Type   string `json:"type"`
}
