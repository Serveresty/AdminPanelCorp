package server

import (
	"AdminPanelCorp/database"
	"AdminPanelCorp/database/access"
	"AdminPanelCorp/database/roleact"
	"AdminPanelCorp/database/useract"
	"AdminPanelCorp/env"
	"AdminPanelCorp/models"
	"AdminPanelCorp/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Функция для POST запроса регистрации
func (db *DataBase) SignUp(c *gin.Context) {
	var user []models.User
	var records [][]string
	var errUsers []models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, elem := range user {
		//Проверка на существование пользователя
		if useract.IsUserRegistered(db.Data, elem.Email, elem.Username) {
			errUsers = append(errUsers, elem)
			continue
		}
		records = append(records, []string{elem.Email, elem.Username})
	}

	data, emailError := useract.CreateUsers(db.Data, records) //Отправка данных на создание пользователей
	if emailError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": emailError})
	}

	errMail := utils.SendEmail(data) //Отправление готовых данных в отправку сообщений на почты
	if errMail != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errMail})
	}

	if len(errUsers) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already registered", "error_data": errUsers})
	}

	if len(records) > 0 {
		c.JSON(http.StatusOK, gin.H{"success": "user has been registered"})
	}

}

// Функция для POST запроса авторизации
func (db *DataBase) SignIn(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Проверка на существование пользователя
	if !useract.IsUserRegistered(db.Data, user.Email, user.Username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user doesn't registered"})
		return
	}
	//Проверка на соответствие паролей в БД с введенным пользователем
	if err := database.CheckPassword(db.Data, user.Email, user.Password); err != nil {
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

	roles, errRoles := roleact.GetUsersRoles(db.Data, user.Id)
	if errRoles != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errRoles})
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

	c.JSON(http.StatusOK, gin.H{"success": "user logged in"})
}

// Функция для POST запроса выхода из сессии
func Logout(c *gin.Context) {
	c.Header("Authorization", "")
	c.JSON(http.StatusOK, gin.H{"success": "user logged out"})
	c.Redirect(http.StatusFound, "/sign-in")
}

func (db *DataBase) AddRoleAccess(c *gin.Context) {
	var accessRoles models.AccessRoles
	if err := c.ShouldBindJSON(&accessRoles); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err1 := access.AddAccessesToRole(db.Data, accessRoles.Role, accessRoles.AccessRoles)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while adding access to role"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "Access to role granted"})
}

func (db *DataBase) DeleteRoleAccess(c *gin.Context) {
	var accessRoles models.AccessRoles
	if err := c.ShouldBindJSON(&accessRoles); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err1 := access.DeleteAccessesFromRole(db.Data, accessRoles.Role, accessRoles.AccessRoles)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while deleting access to role"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "Access to role taken away"})
}
