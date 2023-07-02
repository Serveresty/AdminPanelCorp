package server

import (
	"AdminPanelCorp/database"
	"AdminPanelCorp/env"
	"AdminPanelCorp/models"
	"AdminPanelCorp/utils"
	"fmt"
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

	c.JSON(200, gin.H{"success": "user registered"})
}

// Функция для POST запроса авторизации
func (db *DataBase) Sign_In(c *gin.Context) {
	var user models.User
	var role string
	email := c.PostForm("email")       //Получение из формы email'а
	password := c.PostForm("password") //Получение из формы password'а

	//Проверка на существование пользователя
	if !utils.IsUserRegistered(db.Data, email) {
		c.JSON(401, gin.H{"error": "user doesn't registered"})
		return
	}
	//Проверка на соответствие паролей в БД с введенным пользователем
	if err := utils.CheckPassword(db.Data, email, password); err != nil {
		c.JSON(401, gin.H{"error": "wrong password"})
		return
	}

	getuser := "select user_id, email, username from users_data where users_data.email=$1"
	row, err := db.Data.Query(getuser, email)
	if err != nil {
		fmt.Println("first")
		fmt.Println(err)
		return
	}
	defer row.Close()
	for row.Next() {
		if err := row.Scan(&user.Id, &user.Email, &user.Username); err != nil {
			fmt.Println("second")
			fmt.Println(err)
			return
		}
	}

	getuser_roles := "select roles.role_name from users_data join users_roles on (users_roles.user_id=users_data.user_id) join roles on (roles.role_id=users_roles.role_id) where users_data.email=$1;"
	roww, err := db.Data.Query(getuser_roles, email)
	if err != nil {
		fmt.Println("firsttt")
		fmt.Println(err)
		return
	}
	defer roww.Close()
	for roww.Next() {
		if err := roww.Scan(&role); err != nil {
			fmt.Println("seconddd")
			fmt.Println(err)
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
		c.JSON(500, gin.H{"error": "could not generate token"})
		return
	}

	//Создание куки
	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)

	c.JSON(200, gin.H{"success": "user logged in"})
}

// Функция для POST запроса выхода из сессии
func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true) //Удаление куки
	//c.JSON(200, gin.H{"success": "user logged out"})
	c.Redirect(http.StatusFound, "/sign-in")
}
