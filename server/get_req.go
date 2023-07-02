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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "already authorized"})
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "already authorized"})
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

	all_users_data, err3 := utils.GetAllUsers(db.Data)

	if err3 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while parsing users from db"})
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
