package utils

import (
	"AdminPanelCorp/env"
	"AdminPanelCorp/models"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// Считывание токена
func ParseToken(tokenString string) (claims *models.Claims, err error) {
	my_key := []byte(env.GetEnv("SECRET_KEY"))
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return my_key, err
	})

	if err != nil {
		fmt.Println("first")
		fmt.Println(err)
		return nil, err
	}

	claims, ok := token.Claims.(*models.Claims)

	if !ok {
		fmt.Println("second")
		fmt.Println(err)
		return nil, err
	}

	return claims, nil
}
