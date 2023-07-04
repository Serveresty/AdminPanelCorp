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
	var user []models.User
	var records [][]string
	var err_users []models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, elem := range user {
		//Проверка на существование пользователя
		if utils.IsUserRegistered(db.Data, elem.Email, elem.Username) {
			err_users = append(err_users, elem)
			continue
		}
		records = append(records, []string{elem.Email, elem.Username})
	}

	data, email_error := database.CreateUsers(db.Data, records) //Отправка данных на создание пользователей
	if email_error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": email_error})
	}

	err_mail := utils.Send_Email(data) //Отправление готовых данных в отправку сообщений на почты
	if err_mail != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err_mail})
	}

	if len(err_users) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already registered", "error_data": err_users})
	}

	if len(records) > 0 {
		c.JSON(http.StatusOK, gin.H{"success": "user has been registered"})
	}

}

// Функция для POST запроса авторизации
func (db *DataBase) Sign_In(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Проверка на существование пользователя
	if !utils.IsUserRegistered(db.Data, user.Email, user.Username) {
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

	roles, err_roles := utils.GetUsersRoles(db.Data, user.Email)
	if err_roles != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err_roles})
		return
	}
	user.Role = roles

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

	bearer := "Bearer " + tokenString
	c.Header("Authorization", bearer)

	//Создание куки
	//c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"success": "user logged in"})
}

// Функция для POST запроса выхода из сессии
func Logout(c *gin.Context) {
	c.Header("Authorization", "")
	c.JSON(http.StatusOK, gin.H{"success": "user logged out"})
	c.Redirect(http.StatusFound, "/sign-in")
}
