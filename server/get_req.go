package server

import (
	"AdminPanelCorp/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Главная страница при GET запросе
func (db *DataBase) HomePage(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if len(claims.StandardClaims.Subject) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "access granted"})
}

// Страница авторизации при GET запросе
func (db *DataBase) SignInPage(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "already authorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "access granted"})
}

// Страница регистрации при GET запросе
func (db *DataBase) SignUpPage(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "already authorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "access granted"})
}

// Функция для GET запроса на Админ Панель
func (db *DataBase) AdminPanel(c *gin.Context) {
	var access bool

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}
	token_string := strings.Split(token, " ")[1]

	claims, err2 := utils.ParseToken(token_string)

	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}

	for _, k := range claims.Role {
		if k == "admin" || k == "manager" {
			access = true
		}
	}
	if !access {
		c.JSON(http.StatusForbidden, gin.H{"error": "No rights to access this page"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "access granted"})

	all_users_data, err3 := utils.GetAllUsers(db.Data)

	if err3 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while parsing users from db"})
	}

	c.JSON(http.StatusOK, all_users_data)
}
