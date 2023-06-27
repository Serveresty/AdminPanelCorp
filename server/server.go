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

type User struct {
	Id       int
	Email    string
	Username string
	Role     string
}

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
	var u User
	email := c.PostForm("email")
	username := c.PostForm("username")

	// Проверка на уже зарегистрированного пользователя

	//
	password := utils.Generate_Password() //Генерация пароля

	hash_password, _ := utils.HashPassword(password)                                                //
	queryInsertNewUser := `INSERT INTO "users_data" (email, username, password) VALUES($1, $2, $3)` //Добавляем нового пользователя
	db.Data.MustExec(queryInsertNewUser, email, username, hash_password)                            //

	getuser := "select * from users_data where username = $1" //
	db.Data.Get(&u, getuser, username)                        //Получаем пользователя по уникальному имени

	queryInsertUsersRole := `INSERT INTO "users_roles" (user_id, role_id) VALUES($1, $2)` //
	db.Data.MustExec(queryInsertUsersRole, u.Id, 1)                                       //Добавляем роль пользователю

	data := [][3]string{{email, username, password}}
	utils.Send_Email(data) //Отправка сообщения с данными пользователя
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
