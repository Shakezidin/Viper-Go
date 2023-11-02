package models

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    int    `json:"phone"`
}

//Login Model for users
type Login struct {
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
