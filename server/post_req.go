package server

import (
	"AdminPanelCorp/database"
	"AdminPanelCorp/models"
	"AdminPanelCorp/utils"
	"log"
	"net/http"
	"os"
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
}

// Функция для POST запроса авторизации
func (db *DataBase) Sign_In(c *gin.Context) {
	var user models.User
	email := c.PostForm("email")       //Получение из формы email'а
	password := c.PostForm("password") //Получение из формы password'а

	//Проверка на существование пользователя
	if !utils.IsUserRegistered(db.Data, email) {
		log.Fatal("user doesn't registered")
	}
	//Проверка на соответствие паролей в БД с введенным пользователем
	if err := utils.CheckPassword(db.Data, email, password); err != nil {
		log.Fatal("wrong password")
	}
	getuser := "select users_data.user_id, email, username, roles.role_name from users_data join users_roles on (users_roles.user_id=users_data.user_id) join roles on (roles.role_id=users_roles.role_id) where users_data.email=$1"
	row, err := db.Data.Query(getuser, email)
	if err != nil {
		return
	}
	defer row.Close()
	for row.Next() {
		if err := row.Scan(&user.Id, &user.Email, &user.Username, &user.Role); err != nil {
			return
		}
	}

	expirationTime := time.Now().Add(time.Hour * 24)

	//Клейм для JWT token
	claims := &models.Claims{
		Id:       user.Id,
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	//Создание нового токена с клеймами
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Кодирование токена по ключу
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(500, gin.H{"error": "could not generate token"})
		return
	}

	//Создание куки
	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)

	//c.JSON(200, gin.H{"success": "user logged in"})
	c.Redirect(http.StatusFound, "/")
}

// Функция для POST запроса выхода из сессии
func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true) //Удаление куки
	//c.JSON(200, gin.H{"success": "user logged out"})
	c.Redirect(http.StatusFound, "/sign-in")
}
