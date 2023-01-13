package models

type AuthLoginModel struct {
	Identification string `json:"identification"`
	Password       string `json:"password"`
}
