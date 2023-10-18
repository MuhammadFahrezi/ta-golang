package models

type User struct {
	User_id  int `gorm:"primaryKey"`
	Role_id  int
	Username string
	Email    string
	Password string
}

type Register struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
