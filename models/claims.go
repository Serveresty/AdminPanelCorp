package models

import "github.com/dgrijalva/jwt-go"

//Структура клеймов для JWT
type Claims struct {
	Role []string
	jwt.StandardClaims
}
