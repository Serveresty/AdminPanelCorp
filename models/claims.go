package models

import "github.com/dgrijalva/jwt-go"

//Структура клеймов для JWT
type Claims struct {
	Id       int
	Username string
	Role     string
	jwt.StandardClaims
}
