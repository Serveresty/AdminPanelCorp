package utils

import "golang.org/x/crypto/bcrypt"

// Функция хэширования пароля
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) //Хэширование пароля
	return string(bytes), err
}
