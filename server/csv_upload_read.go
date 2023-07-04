package server

import (
	"AdminPanelCorp/database"
	"AdminPanelCorp/utils"
	"bufio"
	"mime/multipart"
	"net/http"
	"regexp"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
)

// Функция, получающая файл из <input>
func (db *DataBase) UploadUsers(c *gin.Context) {
	var err_users [][]string
	var records [][]string
	file, handler, err := c.Request.FormFile("user_registration")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	defer file.Close()

	content_type := handler.Header.Get("Content-Type")

	if content_type == "text/csv" || content_type == "text/plain" {
		res, err2 := handler.Open()
		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err2})
			return
		}
		result, err3 := readCSVFile(res)
		if err3 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error while reading file"})
		}

		for _, elem := range result {
			//Проверка на существование пользователя
			if utils.IsUserRegistered(db.Data, elem[0], elem[1]) {
				err_users = append(err_users, elem)
				continue
			}
			records = append(records, []string{elem[0], elem[1]})
		}

		data, email_error := database.CreateUsers(db.Data, records) //Отправление данных вида (email, username) в функцию создания пользователей
		if email_error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": email_error})
		}
		err_mail := utils.Send_Email(data) //Отправление готовых данных в отправку сообщений на почты
		if err_mail != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err_mail})
		}

		if len(err_users) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already registered", "error_data": err_users})
		}

		if len(records) > 0 {
			c.JSON(http.StatusOK, gin.H{"success": "user has been registered"})
		}
	} else {
		if content_type == "application/vnd.ms-excel" || content_type == "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
			res, err2 := handler.Open()
			if err2 != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err2})
				return
			}
			result, err3 := readXLSXFile(res)
			if err3 != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Error while reading file"})
			}
			c.JSON(http.StatusOK, gin.H{"message": result})

			for _, elem := range result {
				//Проверка на существование пользователя
				if utils.IsUserRegistered(db.Data, elem[0], elem[1]) {
					err_users = append(err_users, elem)
					continue
				}
				records = append(records, []string{elem[0], elem[1]})
			}

			data, email_error := database.CreateUsers(db.Data, records) //Отправление данных вида (email, username) в функцию создания пользователей
			if email_error != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": email_error})
			}
			err_mail := utils.Send_Email(data) //Отправление готовых данных в отправку сообщений на почты
			if err_mail != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err_mail})
			}
			if len(err_users) > 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "user already registered", "error_data": err_users})
			}

			if len(records) > 0 {
				c.JSON(http.StatusOK, gin.H{"success": "user has been registered"})
			}

		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "wrong file"})
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": "Operation has been completed"})
}

// Функция, читающая файл с email и username
func readCSVFile(fl multipart.File) ([][]string, error) {
	var line []string
	var records [][]string

	fileScanner := bufio.NewScanner(fl)

	//Сканирование файла построчно
	for fileScanner.Scan() {
		words := regexp.MustCompile("[,;\n ]{1}").Split(fileScanner.Text(), -1)
		for _, word := range words {
			if word == "" {
				continue
			}
			line = append(line, word)
		}
		if line == nil {
			continue
		}
		if len(line) > 1 && len(line) < 3 {
			records = append(records, line)
		}
		line = nil
	}

	if err := fileScanner.Err(); err != nil {
		return nil, err
	}
	return records, nil
}

func readXLSXFile(fl multipart.File) ([][]string, error) {
	var result [][]string
	f, err := excelize.OpenReader(fl)
	if err != nil {
		return nil, err
	}
	res := f.GetRows("Лист1")
	for _, elem := range res {
		if strings.ReplaceAll(elem[0], " ", "") != "" && strings.ReplaceAll(elem[1], " ", "") != "" {
			result = append(result, []string{elem[0], elem[1]})
		} else {
			continue
		}
	}
	return result, nil
}
