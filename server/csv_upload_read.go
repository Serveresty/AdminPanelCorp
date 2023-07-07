package server

import (
	"AdminPanelCorp/database/useract"
	"AdminPanelCorp/utils"
	"bufio"
	"mime/multipart"
	"net/http"
	"regexp"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Функция, получающая файл из <input>
func (db *DataBase) UploadUsers(c *gin.Context) {
	file, handler, err := c.Request.FormFile("user_registration")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error db while try to get file"})
		return
	}
	defer file.Close()

	contentType := handler.Header.Get("Content-Type")

	if contentType == "text/csv" || contentType == "text/plain" {
		res, err2 := handler.Open()
		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error db while opening file"})
			return
		}
		result, err3 := readCSVFile(res)
		if err3 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error while reading file"})
		}

		records, errUsers := checkRegistration(db.Data, result)

		data, emailError := useract.CreateUsers(db.Data, records) //Отправление данных вида (email, username) в функцию создания пользователей
		if emailError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": emailError})
		}
		errMail := utils.SendEmail(data) //Отправление готовых данных в отправку сообщений на почты
		if errMail != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errMail})
		}

		if len(errUsers) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already registered", "error_data": errUsers})
		}

		if len(records) > 0 {
			c.JSON(http.StatusOK, gin.H{"success": "user has been registered"})
		}
	} else {
		if contentType == "application/vnd.ms-excel" || contentType == "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
			res, err2 := handler.Open()
			if err2 != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "error db while opening file"})
				return
			}
			result, err3 := readXLSXFile(res)
			if err3 != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Error while reading file"})
			}

			records, errUsers := checkRegistration(db.Data, result)

			data, emailError := useract.CreateUsers(db.Data, records) //Отправление данных вида (email, username) в функцию создания пользователей
			if emailError != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": emailError})
			}
			err_mail := utils.SendEmail(data) //Отправление готовых данных в отправку сообщений на почты
			if err_mail != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err_mail})
			}
			if len(errUsers) > 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "user already registered", "error_data": errUsers})
			}

			if len(records) > 0 {
				c.JSON(http.StatusOK, gin.H{"success": "user has been registered"})
			}

		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported file-type"})
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

func checkRegistration(db *sqlx.DB, data [][]string) ([][]string, [][]string) {
	var errUsers [][]string
	var records [][]string
	for _, elem := range data {
		//Проверка на существование пользователя
		if useract.IsUserRegistered(db, elem[0], elem[1]) {
			errUsers = append(errUsers, elem)
			continue
		}
		records = append(records, []string{elem[0], elem[1]})
	}
	return records, errUsers
}
