package main

import (
	"AdminPanelCorp/database"
	"AdminPanelCorp/env"
	"AdminPanelCorp/models"
	"AdminPanelCorp/requests"
	"AdminPanelCorp/swagger_tests"
	"fmt"
	"log"

	apiclient "AdminPanelCorp/docs/output_swag/client"

	"github.com/gin-gonic/gin"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// @title Admin Panel API
// @version 1.0
// @description API Server

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	go func() {
		transport := httptransport.New(apiclient.DefaultHost, apiclient.DefaultBasePath, nil)
		// create the API client, with the transport
		clientt := apiclient.New(transport, strfmt.Default)
		fmt.Println("--------------Scenario1--------------")
		err1 := swagger_tests.Scenario1(clientt)
		if err1 != nil {
			fmt.Println(err1.Error())
		}
		fmt.Println("--------------Scenario2--------------")
		err2 := swagger_tests.Scenario2(clientt)
		if err2 != nil {
			fmt.Println(err2.Error())
		}
		fmt.Println("--------------Scenario3--------------")
		err3 := swagger_tests.Scenario3(clientt)
		if err3 != nil {
			fmt.Println(err3.Error())
		}
	}()

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
	err1 := database.CreateTable(db)
	if err1 != nil {
		log.Fatal("failed to create tables: ", err.Error())
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	requests.AllRequests(router, db)

	router.Run(":8080")
	return nil
}
