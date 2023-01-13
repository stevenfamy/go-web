package models

type User struct {
	ID           string `json:"id"`
	UserName     string `json:"username"`
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
	CreatedAt    int    `json:"created_at"`
	UpdatedAt    int    `json:"updated_at"`
}
