package requests

import (
	docs "AdminPanelCorp/docs"
	"AdminPanelCorp/server"

	//docs "AdminPanelCorp/docs"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Хэндлеры на запросы
func AllRequests(route *gin.Engine, DB *sqlx.DB) {
	handlerDB := server.DataBase{Data: DB}

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := route.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", server.Helloworld)
		}

		auth := v1.Group("/auth")
		{
			auth.POST("/registration-form", handlerDB.SignUp)
			auth.GET("/sign-up", handlerDB.SignUpPage)
			auth.GET("/sign-in", handlerDB.SignInPage)
			auth.POST("/login-form", handlerDB.SignIn)
			auth.POST("/logout-form", server.Logout)
		}

		pgs := v1.Group("/page")
		{
			pgs.GET("/homepage", handlerDB.HomePage)
			pgs.GET("/admin", handlerDB.AdminPanel)
		}
		v1.POST("/edit-user", handlerDB.EditUser)
		v1.POST("/add-role", handlerDB.AddRole)
		v1.POST("/delete-role", handlerDB.DeleteRole)
		v1.POST("/delete-user", handlerDB.DeleteUser)
		v1.POST("/add-role-access", handlerDB.AddRoleAccess)
		v1.POST("/delete-role-access", handlerDB.DeleteRoleAccess)
		v1.POST("/upload", handlerDB.UploadUsers)
	}
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
