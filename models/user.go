package models

//Структура пользователя
type User struct {
	Id       int      `json:"id"`
	Email    string   `json:"email"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Role     []string `json:"role"`
}

type RegisterUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
