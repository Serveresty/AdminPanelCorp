package server

import (
	"AdminPanelCorp/database"
	"AdminPanelCorp/utils"
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
)

// Функция, получающая файл из <input>
func (db *DataBase) UploadUsers(c *gin.Context) {
	fileObj, err := c.FormFile("filename") //Получение csv файла из html
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}
	filePath := fmt.Sprintf("./%s", fileObj.Filename) //Задавание пути csv файла
	c.SaveUploadedFile(fileObj, filePath)             //Сохранение csv файла

	records := readCSVFile(filePath)               //Считывание csv файла
	data := database.CreateUsers(db.Data, records) //Отправление данных вида (email, username) в функцию создания пользователей
	utils.Send_Email(data)                         //Отправление готовых данных в отправку сообщений на почты

	eerr := os.Remove(filePath) //Удаление csv файла из Path
	if eerr != nil {
		log.Fatal(eerr)
	}
	c.JSON(http.StatusOK, gin.H{"success": "all users has been added"})
}

// Функция, читающая файл с email и username
func readCSVFile(filePath string) [][]string {
	var line []string
	var records [][]string
	file, err := os.Open(filePath) //Открытие файла
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	//Сканирование файла построчно
	for fileScanner.Scan() {
		words := regexp.MustCompile("[,;\n]{1}").Split(fileScanner.Text(), -1)
		for _, word := range words {
			line = append(line, word)
		}
		records = append(records, line)
		line = nil
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	return records
}
