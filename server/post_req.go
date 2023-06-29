package server

import (
	"AdminPanelCorp/models"
	"AdminPanelCorp/utils"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Функция для POST запроса регистрации
func (db *DataBase) Sign_Up(c *gin.Context) {
	var id int
	var slice []string
	var data [][]string
	emails := c.PostFormArray("email")
	usernames := c.PostFormArray("username")

	queryInsertNewUser := `INSERT INTO "users_data" (email, username, password) VALUES($1, $2, $3)`
	getuser := "select user_id from users_data where username = $1"
	queryInsertUsersRole := `INSERT INTO "users_roles" (user_id, role_id) VALUES($1, $2)`

	for i, r := range emails {
		password := utils.Generate_Password()                                //Генерация пароля
		hash_password, _ := utils.HashPassword(password)                     //Хэш пароля
		db.Data.MustExec(queryInsertNewUser, r, usernames[i], hash_password) //Добавление в базу нового пользователя
		db.Data.Get(&id, getuser, usernames[i])                              //Получить id пользователя по username
		db.Data.MustExec(queryInsertUsersRole, &id, 1)                       //Присвоение роли user по id пользователя
		slice = append(slice, r, usernames[i], password)                     //Объединение почты, username и пароля в слайс
		data = append(data, slice)                                           //Добавление слайса в слайс для передачи в функцию отправки письма на почту
		slice = nil
	}

	utils.Send_Email(data) //Отправка сообщения на почту с данными пользователя
}

// Функция для POST запроса авторизации
func (db *DataBase) Sign_In(c *gin.Context) {
	var user models.User
	email := c.PostForm("email")
	password := c.PostForm("password")

	if !utils.IsUserRegistered(db.Data, email) {
		log.Fatal("user doesn't registered")
	}
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

	claims := &models.Claims{
		Id:       user.Id,
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(500, gin.H{"error": "could not generate token"})
		return
	}
	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": "user logged in"})
}

// Функция для POST запроса выхода из сессии
func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": "user logged out"})
}
