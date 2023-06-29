package main

import (
	"AdminPanelCorp/database"
	"AdminPanelCorp/models"
	"AdminPanelCorp/requests"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	//Подключаемся к БД(через структуру Config передаем данные для подключения)
	db, err := database.DB_Init(&models.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "abegah54",
		DBName:   "postgres",
		SSLMode:  "disable",
	})

	defer db.Close()

	if err != nil {
		log.Fatal("failed to initialize db: ", err.Error())
	}

	//Создаем таблицы с пользователями, ролями и их соответствием
	err1, err2, err3 := database.Create_Table(db)
	if err1 != nil {
		log.Fatal("failed to create users_data table: ", err1.Error())
	}
	if err2 != nil {
		log.Fatal("failed to create roles table: ", err2.Error())
	}
	if err3 != nil {
		log.Fatal("failed to create users_roles table: ", err3.Error())
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	requests.All_requests(router, db)

	router.Run(":8080")
}
