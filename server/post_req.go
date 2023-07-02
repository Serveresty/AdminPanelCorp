package server

import (
	"AdminPanelCorp/database"
	"AdminPanelCorp/env"
	"AdminPanelCorp/models"
	"AdminPanelCorp/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Функция для POST запроса регистрации
func (db *DataBase) Sign_Up(c *gin.Context) {
	var slice []string
	var records [][]string
	emails := c.PostFormArray("email")
	usernames := c.PostFormArray("username")

	for i, r := range emails {
		slice = append(slice, r, usernames[i]) //Объединение почты, username и пароля в слайс
		records = append(records, slice)       //Добавление слайса в слайс для передачи в функцию отправки письма на почту
		slice = nil
	}

	data := database.CreateUsers(db.Data, records) //Отправка данных на создание пользователей

	utils.Send_Email(data) //Отправка сообщения на почту с данными пользователя

	c.JSON(http.StatusOK, gin.H{"success": "user has been registered"})
}

// Функция для POST запроса авторизации
func (db *DataBase) Sign_In(c *gin.Context) {
	var user models.User
	var role string

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Проверка на существование пользователя
	if !utils.IsUserRegistered(db.Data, user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user doesn't registered"})
		return
	}
	//Проверка на соответствие паролей в БД с введенным пользователем
	if err := utils.CheckPassword(db.Data, user.Email, user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong password"})
		return
	}

	getuser := "select user_id, username from users_data where users_data.email=$1"
	row, err := db.Data.Query(getuser, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	defer row.Close()
	for row.Next() {
		if err := row.Scan(&user.Id, &user.Username); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	getuser_roles := "select roles.role_name from users_data join users_roles on (users_roles.user_id=users_data.user_id) join roles on (roles.role_id=users_roles.role_id) where users_data.email=$1;"
	roww, err := db.Data.Query(getuser_roles, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	defer roww.Close()
	for roww.Next() {
		if err := roww.Scan(&role); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		user.Role = append(user.Role, role)
		role = ""
	}

	expirationTime := time.Now().Add(time.Hour * 24)

	//Клейм для JWT token
	claims := &models.Claims{
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	//Создание нового токена с клеймами
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Кодирование токена по ключу
	tokenString, err := token.SignedString([]byte(env.GetEnv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	//Создание куки
	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"success": "user logged in"})
}

// Функция для POST запроса выхода из сессии
func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true) //Удаление куки
	c.JSON(http.StatusOK, gin.H{"success": "user logged out"})
	c.Redirect(http.StatusFound, "/sign-in")
}
