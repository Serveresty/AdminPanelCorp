package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Функция добавления роли менеджера админом
func (db *DataBase) AddManagerRole(c *gin.Context) {
	userId := c.PostForm("userId")
	queryInsertUsersRole := `INSERT INTO users_roles (user_id, role_id) SELECT users_data.user_id, roles.role_id FROM users_data, roles WHERE users_data.user_id=$1 and roles.role_name=$2;`
	db.Data.MustExec(queryInsertUsersRole, &userId, "manager")
	c.Redirect(http.StatusFound, "/admin")
}

// Функция удаления роли менеджера админом
func (db *DataBase) DeleteManagerRole(c *gin.Context) {
	userId := c.PostForm("userId")
	queryInsertUsersRole := `DELETE FROM users_roles WHERE user_id=$1 and role_id=(SELECT role_id FROM roles WHERE role_name = $2)`
	db.Data.MustExec(queryInsertUsersRole, &userId, "manager")
	c.Redirect(http.StatusFound, "/admin")
}

// Функция удаления пользователя админом
func (db *DataBase) DeleteUser(c *gin.Context) {
	userId := c.PostForm("userId")
	queryInsertUsersRole := `DELETE FROM users_data WHERE user_id=$1`
	db.Data.MustExec(queryInsertUsersRole, &userId)
	c.Redirect(http.StatusFound, "/admin")
}
