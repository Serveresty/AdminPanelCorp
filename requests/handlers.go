package requests

import (
	"AdminPanelCorp/server"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Хэндлеры на запросы
func AllRequests(route *gin.Engine, DB *sqlx.DB) {
	handler_db := server.DataBase{Data: DB}
	route.GET("/", handler_db.HomePage)
	route.GET("/sign-up", handler_db.SignUpPage)
	route.GET("/sign-in", handler_db.SignInPage)
	route.GET("/admin", handler_db.AdminPanel)
	route.POST("/login-form", handler_db.SignIn)
	route.POST("/registration-form", handler_db.SignUp)
	route.POST("/logout-form", server.Logout)
	route.POST("/edit-user", handler_db.EditUser)
	route.POST("/add-role", handler_db.AddRole)
	route.POST("/delete-role", handler_db.DeleteRole)
	route.POST("/delete-user", handler_db.DeleteUser)
	route.POST("/upload", handler_db.UploadUsers)
}
