package server

import (
	"AdminPanelCorp/database"
	"AdminPanelCorp/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DataBase struct {
	Data *sqlx.DB
}

//type User struct {
//	Id       int
//	Email    string
//	Username string
//	Role     string
//}

func (db *DataBase) Panel(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"panel.html",
		gin.H{},
	)
}

func (db *DataBase) Sign_In_Page(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"sign_in.html",
		gin.H{},
	)
}

func (db *DataBase) Sign_Up_Page(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"sign_up.html",
		gin.H{},
	)
}

func (db *DataBase) Sign_Up(c *gin.Context) {
	var id int
	var slice []string
	var data [][]string
	emails := c.PostFormArray("email")
	usernames := c.PostFormArray("username")

	queryInsertNewUser := `INSERT INTO "users_data" (email, username, password) VALUES($1, $2, $3)`
	getuser := "select id from users_data where username = $1"
	queryInsertUsersRole := `INSERT INTO "users_roles" (user_id, role_id) VALUES($1, $2)`

	for i, r := range emails {
		password := utils.Generate_Password()                                //Генерация пароля
		hash_password, _ := utils.HashPassword(password)                     //Хэш пароля
		db.Data.MustExec(queryInsertNewUser, r, usernames[i], hash_password) //Добавление в базу нового пользователя
		db.Data.Get(&id, getuser, usernames[i])                              //Получить id пользователя по username
		db.Data.MustExec(queryInsertUsersRole, id, 1)                        //Присвоение роли user по id пользователя
		slice = append(slice, r, usernames[i], password)                     //Объединение почты, username и пароля в слайс
		data = append(data, slice)                                           //Добавление слайса в слайс для передачи в функцию отправки письма на почту
		slice = nil
	}

	utils.Send_Email(data) //Отправка сообщения на почту с данными пользователя
}

func (db *DataBase) Sign_In(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	if !database.IsUserRegistered(db.Data, email) {
		log.Fatal("user doesn't registered")
	}
	if err := database.CheckPassword(db.Data, email, password); err != nil {
		log.Fatal("wrong password")
	}
}
