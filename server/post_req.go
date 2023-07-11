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

// @BasePath /api/v1

// PingExample godoc
// @Summary SignUp
// @Tags auth
// @Description create account
// @Accept json
// @Produce json
// @Param input body []models.RegisterUser true "account info"
// @Success 200 {string} string "success"
// @Failure 400 {string} string "error"
// @Failure 500 {string} [][]string "error"
// @Router /auth/registration-form [post]
func (db *DataBase) SignUp(c *gin.Context) {
	var user []models.RegisterUser
	var records [][]string
	var errUsers []models.RegisterUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input body"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": emailError})
	}

	errMail := utils.SendEmail(data) //Отправление готовых данных в отправку сообщений на почты
	if errMail != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMail})
	}

	if len(errUsers) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already registered", "error_data": errUsers})
	}

	if len(records) > 0 {
		c.JSON(http.StatusOK, gin.H{"success": "user has been registered"})
	}

}

// @BasePath /api/v1

// PingExample godoc
// @Summary SignIn
// @Tags auth
// @Description login
// @Accept json
// @Produce json
// @Param input body models.User true "account info"
// @Success 200 {string} string "success"
// @Failure 400 {string} string "error"
// @Failure 500 {string} string "error"
// @Router /auth/login-form [post]
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

// @BasePath /api/v1

// PingExample godoc
// @Summary Logout
// @Security ApiKeyAuth
// @Tags auth
// @Description logout
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
// @Router /auth/logout-form [post]
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

	var isAccess bool

	claims, errStr := parseInfoFromToken(c)
	if errStr != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}

	for _, k := range claims.Role {
		if k == "admin" {
			isAccess = true
		}
	}

	if !isAccess {
		c.JSON(http.StatusForbidden, gin.H{"error": "No rights to action"})
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

	var isAccess bool

	claims, errStr := parseInfoFromToken(c)
	if errStr != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}

	for _, k := range claims.Role {
		if k == "admin" {
			isAccess = true
		}
	}

	if !isAccess {
		c.JSON(http.StatusForbidden, gin.H{"error": "No rights to action"})
		return
	}

	err1 := access.DeleteAccessesFromRole(db.Data, accessRoles.Role, accessRoles.AccessRoles)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while deleting access to role"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "Access to role taken away"})
}
