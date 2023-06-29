package server

import (
	"AdminPanelCorp/database"
	"AdminPanelCorp/models"
	"AdminPanelCorp/utils"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Структура с указателем на БД
type DataBase struct {
	Data *sqlx.DB
}

// Главная страница при GET запросе
func (db *DataBase) Home_Page(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	if claims.Role != "" {

	}

	c.HTML(
		http.StatusOK,
		"home_page.html",
		gin.H{},
	)
}

// Страница авторизации при GET запросе
func (db *DataBase) Sign_In_Page(c *gin.Context) {
	_, err := c.Cookie("token")

	if err == nil {
		c.JSON(401, gin.H{"error": "already authorized"})
		return
	}
	c.HTML(
		http.StatusOK,
		"sign_in.html",
		gin.H{},
	)
}

// Страница регистрации при GET запросе
func (db *DataBase) Sign_Up_Page(c *gin.Context) {
	_, err := c.Cookie("token")

	if err == nil {
		c.JSON(401, gin.H{"error": "already authorized"})
		return
	}
	c.HTML(
		http.StatusOK,
		"sign_up.html",
		gin.H{},
	)
}

// Функция для POST запроса регистрации
func (db *DataBase) Sign_Up(c *gin.Context) {
	var id int
	var slice []string
	var data [][]string
	emails := c.PostFormArray("email")
	usernames := c.PostFormArray("username")

	queryInsertNewUser := `INSERT INTO "users_data" (email, username, password) VALUES($1, $2, $3)`
	getuser := "select user_id from users_data where username = $1"
	queryInsertUsersRole := `INSERT INTO "users_roles" (user_id, role_id) VALUES($1, $2)`

	for i, r := range emails {
		password := utils.Generate_Password()                                //Генерация пароля
		hash_password, _ := utils.HashPassword(password)                     //Хэш пароля
		db.Data.MustExec(queryInsertNewUser, r, usernames[i], hash_password) //Добавление в базу нового пользователя
		db.Data.Get(&id, getuser, usernames[i])                              //Получить id пользователя по username
		db.Data.MustExec(queryInsertUsersRole, &id, 1)                       //Присвоение роли user по id пользователя
		slice = append(slice, r, usernames[i], password)                     //Объединение почты, username и пароля в слайс
		data = append(data, slice)                                           //Добавление слайса в слайс для передачи в функцию отправки письма на почту
		slice = nil
	}

	utils.Send_Email(data) //Отправка сообщения на почту с данными пользователя
}

// Функция для POST запроса авторизации
func (db *DataBase) Sign_In(c *gin.Context) {
	var user models.User
	email := c.PostForm("email")
	password := c.PostForm("password")

	if !database.IsUserRegistered(db.Data, email) {
		log.Fatal("user doesn't registered")
	}
	if err := database.CheckPassword(db.Data, email, password); err != nil {
		log.Fatal("wrong password")
	}
	getuser := "select users_data.user_id, email, username, roles.role_name from users_data join users_roles on (users_roles.user_id=users_data.user_id) join roles on (roles.role_id=users_roles.role_id) where users_data.email=$1"
	row, err := db.Data.Query(getuser, email)
	if err != nil {
		return
	}
	defer row.Close()
	for row.Next() {
		if err := row.Scan(&user.Id, &user.Email, &user.Username, &user.Role); err != nil {
			return
		}
	}

	expirationTime := time.Now().Add(time.Hour * 24)

	claims := &models.Claims{
		Id:       user.Id,
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(500, gin.H{"error": "could not generate token"})
		return
	}
	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": "user logged in"})
}

// Функция для POST запроса выхода из сессии
func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": "user logged out"})
}

// Функция для GET запроса на Админ Панель
func (db *DataBase) Admin_Panel(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "admin" && claims.Role != "manager" {
		c.JSON(401, gin.H{"error": "no permissions to access this page"})
		return
	}

	all_users_data, err := database.GetAllUsers(db.Data)

	c.HTML(
		http.StatusOK,
		"panel.html",
		gin.H{
			"title": "Admin Panel",
			"data":  all_users_data,
		},
	)
}

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

// Функция добавления роли менеджера админом
func (db *DataBase) AddManagerRole(c *gin.Context) {
	userId := c.PostForm("userId")
	queryInsertUsersRole := `UPDATE users_roles SET role_id = (SELECT role_id FROM roles WHERE role_name = $1) where user_id = $2;`
	db.Data.MustExec(queryInsertUsersRole, "manager", &userId)
}

// Функция удаления роли менеджера админом
func (db *DataBase) DeleteManagerRole(c *gin.Context) {
	userId := c.PostForm("userId")
	queryInsertUsersRole := `UPDATE users_roles SET role_id = (SELECT role_id FROM roles WHERE role_name = $1) where user_id = $2;`
	db.Data.MustExec(queryInsertUsersRole, "user", &userId)
}

// Функция удаления пользователя админом
func (db *DataBase) DeleteUser(c *gin.Context) {
	userId := c.PostForm("userId")
	queryInsertUsersRole := `DELETE FROM users_data WHERE user_id=$1`
	db.Data.MustExec(queryInsertUsersRole, &userId)
}
