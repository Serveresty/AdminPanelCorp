package main

import (
	"AdminPanelCorp/database"
	"AdminPanelCorp/env"
	"AdminPanelCorp/models"
	"AdminPanelCorp/requests"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	//Подключаемся к БД(через структуру Config передаем данные для подключения)
	db, err := database.DBInit(&models.Config{
		Host:     env.GetEnv("Host"),
		Port:     env.GetEnv("Port"),
		Username: env.GetEnv("Username"),
		Password: env.GetEnv("Password"),
		DBName:   env.GetEnv("DBName"),
		SSLMode:  env.GetEnv("SSLMode"),
	})

	defer db.Close()

	if err != nil {
		log.Fatal("failed to initialize db: ", err.Error())
	}

	//Создаем таблицы с пользователями, ролями и их соответствием
	err1, err2, err3 := database.CreateTable(db)
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
	requests.AllRequests(router, db)

	router.Run(":8080")
	return nil
}
