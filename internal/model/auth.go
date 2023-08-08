package model

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type WpUser struct {
	UserLogin   string `json:"user_login"`
	UserEmail   string `json:"user_email"`
	DisplayName string `json:"display_name"`
}
