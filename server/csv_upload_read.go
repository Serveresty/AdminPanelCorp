package server

import (
	"AdminPanelCorp/database"
	"AdminPanelCorp/utils"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

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

	records := ReadCSVFile(filePath)               //Считывание csv файла
	data := database.CreateUsers(db.Data, records) //Отправление данных вида (email, username) в функцию создания пользователей
	utils.Send_Email(data)                         //Отправление готовых данных в отправку сообщений на почты
	eerr := os.Remove(filePath)                    //Удаление csv файла из Path
	if eerr != nil {
		log.Fatal(eerr)
	}
	c.Redirect(http.StatusFound, "/admin")
}

// Функция, читающая CSV
func ReadCSVFile(filePath string) [][]string {
	f, err := os.Open(filePath) //Открытие csv файла
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll() //Считывание csv файла
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records //Возврат массива из структур с данными пользователей
}
