package main

import (
	"AdminPanelCorp/database"
	"AdminPanelCorp/env"
	"AdminPanelCorp/models"
	"AdminPanelCorp/requests"
	"fmt"
	"log"

	apiclient "AdminPanelCorp/docs/output_swag/client"
	"AdminPanelCorp/docs/output_swag/client/example"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/gin-gonic/gin"
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

// @host localhost:80
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	transport := httptransport.New("localhost:80", apiclient.DefaultBasePath, nil)

	// create the API client, with the transport
	clientt := apiclient.New(transport, strfmt.Default)

	resp, err := clientt.Example.GetExampleHelloworld(example.NewGetExampleHelloworldParams())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp)

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
