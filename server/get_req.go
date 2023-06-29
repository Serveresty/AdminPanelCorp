package server

import (
	"AdminPanelCorp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Главная страница при GET запросе
func (db *DataBase) Home_Page(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	if claims.Role != "" {

	}

	c.HTML(
		http.StatusOK,
		"home_page.html",
		gin.H{},
	)
}

// Страница авторизации при GET запросе
func (db *DataBase) Sign_In_Page(c *gin.Context) {
	_, err := c.Cookie("token")

	if err == nil {
		c.JSON(401, gin.H{"error": "already authorized"})
		return
	}
	c.HTML(
		http.StatusOK,
		"sign_in.html",
		gin.H{},
	)
}

// Страница регистрации при GET запросе
func (db *DataBase) Sign_Up_Page(c *gin.Context) {
	_, err := c.Cookie("token")

	if err == nil {
		c.JSON(401, gin.H{"error": "already authorized"})
		return
	}
	c.HTML(
		http.StatusOK,
		"sign_up.html",
		gin.H{},
	)
}

// Функция для GET запроса на Админ Панель
func (db *DataBase) Admin_Panel(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "admin" && claims.Role != "manager" {
		c.JSON(401, gin.H{"error": "no permissions to access this page"})
		return
	}

	all_users_data, err := utils.GetAllUsers(db.Data)

	c.HTML(
		http.StatusOK,
		"panel.html",
		gin.H{
			"title": "Admin Panel",
			"data":  all_users_data,
		},
	)
}
