package server

import (
	"AdminPanelCorp/database/useract"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Главная страница при GET запросе
func (db *DataBase) HomePage(c *gin.Context) {
	claims, errStr := parseInfoFromToken(c)
	if errStr != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}

	if len(claims.StandardClaims.Subject) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "access granted"})
}

// Страница авторизации при GET запросе

// @BasePath /api/v1

// PingExample godoc
// @Summary SignIn
// @Tags auth
// @Description login
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
// @Failure 400 {string} string "error"
// @Router /auth/sign-in [get]
func (db *DataBase) SignInPage(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "already authorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "access granted"})
}

// Страница регистрации при GET запросе

// @BasePath /api/v1

// PingExample godoc
// @Summary SignUp
// @Tags auth
// @Description create account
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
// @Failure 400 {string} string "error"
// @Router /auth/sign-up [get]
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

	claims, errStr := parseInfoFromToken(c)
	if errStr != "" {
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

	allUsersData, err3 := useract.GetAllUsers(db.Data)

	if err3 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while parsing users from db"})
	}

	c.JSON(http.StatusOK, allUsersData)
}
