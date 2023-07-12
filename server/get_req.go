package server

import (
	"AdminPanelCorp/database/useract"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// Главная страница при GET запросе

// @BasePath /api/v1

// PingExample godoc
// @Summary ping page
// @Security ApiKeyAuth
// @Schemes
// @Description home page
// @Tags page
// @Accept json
// @Produce json
// @Success 202 {string} Access granted
// @Failure 400,401 {string} Access denied
// @Header 200,400,default {string} Authorization "Authorization"
// @Router /page/homepage [get]
func (db *DataBase) HomePage(c *gin.Context) {
	claims, errStr := parseInfoFromToken(c)
	if errStr != "" {
		c.JSON(http.StatusBadRequest, "unauthorized")
		return
	}

	if len(claims.StandardClaims.Subject) == 0 {
		c.JSON(http.StatusUnauthorized, "unauthorized")
	}

	c.JSON(http.StatusAccepted, "access granted")
}

// PingExample godoc
// @Summary ping auth
// @Schemes
// @Description SignIn page
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {string} AccessGranted
// @Failure 400 {string} AccessDenied
// @Router /auth/sign-in [get]
func (db *DataBase) SignInPage(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token != "" {
		c.JSON(http.StatusBadRequest, "already authorized")
		return
	}
	c.JSON(http.StatusOK, "access granted")
}

// Страница регистрации при GET запросе

// PingExample godoc
// @Summary ping auth
// @Schemes
// @Description SignUp page
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {string} AccessGranted
// @Failure 400 {string} AccessDenied
// @Router /auth/sign-up [get]
func (db *DataBase) SignUpPage(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token != "" {
		c.JSON(http.StatusBadRequest, "already authorized")
		return
	}
	c.JSON(http.StatusOK, "access granted")
}

// Функция для GET запроса на Админ Панель
// @Security ApiKeyAuth
// PingExample godoc
// @Summary ping page
// @Schemes
// @Description Admin page
// @Tags page
// @Accept json
// @Produce json
// @Success 200 {string} Access granted
// @Failure 400,403,500 {string} Access denied
// @Header 200,400,403,500,default {string} Authorization "Authorization"
// @Router /page/admin [get]
func (db *DataBase) AdminPanel(c *gin.Context) {
	var access bool

	claims, errStr := parseInfoFromToken(c)
	if errStr != "" {
		c.JSON(http.StatusBadRequest, "unauthorized")
		return
	}

	for _, k := range claims.Role {
		if k == "admin" || k == "manager" {
			access = true
		}
	}

	if !access {
		c.JSON(http.StatusForbidden, "No rights to access this page")
		return
	}

	c.JSON(http.StatusOK, "access granted")

	allUsersData, err3 := useract.GetAllUsers(db.Data)

	if err3 != nil {
		c.JSON(http.StatusInternalServerError, "error while parsing users from db")
	}

	c.JSON(http.StatusOK, allUsersData)
}
