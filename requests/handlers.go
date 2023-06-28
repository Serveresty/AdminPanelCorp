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
	route.GET("/", handler_db.Panel)
	route.GET("/sign-up", handler_db.Sign_Up_Page)
	route.GET("/sign-in", handler_db.Sign_In_Page)
	route.POST("/login-form", handler_db.Sign_In)
	route.POST("/registration-form", handler_db.Sign_Up)
	route.POST("/logout-form", server.Logout)
}
