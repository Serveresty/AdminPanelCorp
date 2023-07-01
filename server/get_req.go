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
		c.JSON(401, gin.H{"error": "unauthorizedH1"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorizedH2"})
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
	_, err := c.Cookie("token") //Проверка на существование куки

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
	_, err := c.Cookie("token") //Проверка на существование куки

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
	cookie, err1 := c.Cookie("token") //Проверка на существование куки
	if err1 != nil {
		c.JSON(401, gin.H{"error": "unauthorized1"})
		return
	}

	claims, err2 := utils.ParseToken(cookie)

	if err2 != nil {
		c.JSON(401, gin.H{"error": "unauthorized2"})
		return
	}

	if claims.Role != "admin" && claims.Role != "manager" {
		c.JSON(401, gin.H{"error": "no permissions to access this page"})
		return
	}

	all_users_data, err3 := utils.GetAllUsers(db.Data)

	if err3 != nil {
		c.JSON(401, gin.H{"error": "error with parsing users from"})
	}

	c.HTML(
		http.StatusOK,
		"panel.html",
		gin.H{
			"title": "Admin Panel",
			"data":  all_users_data,
		},
	)
}
