package requests

import (
	"AdminPanelCorp/server"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func All_requests(route *gin.Engine, DB *sqlx.DB) {
	handler_db := server.DataBase{Data: DB}
	route.LoadHTMLGlob("web/templates/*")
	route.Static("/web", "./web/")
	route.GET("/", handler_db.Home_Page)
	route.GET("/sign-up", handler_db.Sign_Up_Page)
	route.GET("/sign-in", handler_db.Sign_In_Page)
	route.GET("/admin", handler_db.Admin_Panel)
	route.POST("/login-form", handler_db.Sign_In)
	route.POST("/registration-form", handler_db.Sign_Up)
	route.POST("/logout-form", server.Logout)
	route.POST("/add-manager", handler_db.AddManagerRole)
	route.POST("/delete-manager", handler_db.DeleteManagerRole)
	route.POST("/delete-user", handler_db.DeleteUser)
	route.POST("/upload", handler_db.UploadUsers)
}
