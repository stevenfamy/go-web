package models

type AuthLogin struct {
	Identification string `json:"identification"`
	Password       string `json:"password"`
}
