package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Функция добавления роли менеджера админом
func (db *DataBase) AddManagerRole(c *gin.Context) {
	userId := c.PostForm("userId")
	queryInsertUsersRole := `UPDATE users_roles SET role_id = (SELECT role_id FROM roles WHERE role_name = $1) where user_id = $2;`
	db.Data.MustExec(queryInsertUsersRole, "manager", &userId)
	c.Redirect(http.StatusFound, "/admin")
}

// Функция удаления роли менеджера админом
func (db *DataBase) DeleteManagerRole(c *gin.Context) {
	userId := c.PostForm("userId")
	queryInsertUsersRole := `UPDATE users_roles SET role_id = (SELECT role_id FROM roles WHERE role_name = $1) where user_id = $2;`
	db.Data.MustExec(queryInsertUsersRole, "user", &userId)
	c.Redirect(http.StatusFound, "/admin")
}

// Функция удаления пользователя админом
func (db *DataBase) DeleteUser(c *gin.Context) {
	userId := c.PostForm("userId")
	queryInsertUsersRole := `DELETE FROM users_data WHERE user_id=$1`
	db.Data.MustExec(queryInsertUsersRole, &userId)
	c.Redirect(http.StatusFound, "/admin")
}
