package main

import (
	"AdminPanelCorp/database"
	"AdminPanelCorp/requests"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.DB_Init(database.Config{
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

	database.Create_Table(db)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	requests.All_requests(router, db)

	router.Run(":8080")
}
