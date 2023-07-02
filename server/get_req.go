package server

import (
	"AdminPanelCorp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Главная страница при GET запросе
func (db *DataBase) Home_Page(c *gin.Context) {
	cookie, err := c.Cookie("token") //Проверка на существование куки
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

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
func (db *DataBase) Sign_In_Page(c *gin.Context) {
	_, err := c.Cookie("token") //Проверка на существование куки

	if err == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "already authorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "access granted"})
}

// Страница регистрации при GET запросе
func (db *DataBase) Sign_Up_Page(c *gin.Context) {
	_, err := c.Cookie("token") //Проверка на существование куки

	if err == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "already authorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "access granted"})
}

// Функция для GET запроса на Админ Панель
func (db *DataBase) Admin_Panel(c *gin.Context) {
	var access bool = false
	cookie, err1 := c.Cookie("token") //Проверка на существование куки
	if err1 != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	claims, err2 := utils.ParseToken(cookie)

	if err2 != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	//fmt.Println(claims.Role)
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
