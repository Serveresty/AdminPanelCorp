package server

import (
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
	fileObj, err := c.FormFile("filename")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}
	filePath := fmt.Sprintf("./%s", fileObj.Filename)
	c.SaveUploadedFile(fileObj, filePath)

	records := ReadCSVFile(filePath)

	var id int

	queryInsertNewUser := `INSERT INTO "users_data" (email, username, password) VALUES($1, $2, $3)`
	getuser := "select user_id from users_data where username = $1"
	queryInsertUsersRole := `INSERT INTO "users_roles" (user_id, role_id) VALUES($1, $2)`

	for i, element := range records {
		password := utils.Generate_Password()                                       //Генерация пароля
		hash_password, _ := utils.HashPassword(password)                            //Хэш пароля
		db.Data.MustExec(queryInsertNewUser, element[0], element[1], hash_password) //Добавление в базу нового пользователя
		db.Data.Get(&id, getuser, element[1])                                       //Получить id пользователя по username
		db.Data.MustExec(queryInsertUsersRole, &id, 1)                              //Присвоение роли user по id пользователя
		records[i] = append(records[i], password)
	}
	utils.Send_Email(records)
	eerr := os.Remove(filePath)
	if eerr != nil {
		log.Fatal(eerr)
	}
}

// Функция, читающая CSV
func ReadCSVFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
