package requests

import (
	"AdminPanelCorp/server"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Хэндлеры на запросы
func AllRequests(route *gin.Engine, DB *sqlx.DB) {
	handlerDB := server.DataBase{Data: DB}
	route.GET("/", handlerDB.HomePage)
	route.GET("/sign-up", handlerDB.SignUpPage)
	route.GET("/sign-in", handlerDB.SignInPage)
	route.GET("/admin", handlerDB.AdminPanel)
	route.POST("/login-form", handlerDB.SignIn)
	route.POST("/registration-form", handlerDB.SignUp)
	route.POST("/logout-form", server.Logout)
	route.POST("/edit-user", handlerDB.EditUser)
	route.POST("/add-role", handlerDB.AddRole)
	route.POST("/delete-role", handlerDB.DeleteRole)
	route.POST("/delete-user", handlerDB.DeleteUser)
	route.POST("/add-role-access", handlerDB.AddRoleAccess)
	route.POST("/delete-role-access", handlerDB.DeleteRoleAccess)
	route.POST("/upload", handlerDB.UploadUsers)
}
