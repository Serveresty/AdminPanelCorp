package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Id       int
	Username string
	Role     string
	jwt.StandardClaims
}
