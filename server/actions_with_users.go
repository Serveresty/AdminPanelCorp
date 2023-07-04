package server

import (
	"AdminPanelCorp/models"
	"AdminPanelCorp/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func (db *DataBase) EditUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if strings.ReplaceAll(user.Email, " ", "") != "" && strings.ReplaceAll(user.Username, " ", "") != "" {
		if utils.IsEmailValid(user.Email) {
			queryInsertNewUsersData := `UPDATE users_data SET email = $2, username = $3 WHERE user_id=$1`
			db.Data.MustExec(queryInsertNewUsersData, &user.Id, &user.Email, &user.Username)
			c.JSON(http.StatusOK, gin.H{"success": "Username and Email has been changed"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not valid email"})
		}
		return
	}
	if strings.ReplaceAll(user.Email, " ", "") != "" && strings.ReplaceAll(user.Username, " ", "") == "" {
		if utils.IsEmailValid(user.Email) {
			queryInsertNewUsersData := `UPDATE users_data SET email = $1 WHERE user_id=$2`
			db.Data.MustExec(queryInsertNewUsersData, &user.Email, &user.Id)
			c.JSON(http.StatusOK, gin.H{"success": "Email has been changed"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not valid email"})
		}
		return
	}
	if strings.ReplaceAll(user.Email, " ", "") == "" && strings.ReplaceAll(user.Username, " ", "") != "" {
		queryInsertNewUsersData := `UPDATE users_data SET username = $1 WHERE user_id=$2`
		db.Data.MustExec(queryInsertNewUsersData, &user.Username, &user.Id)
		c.JSON(http.StatusOK, gin.H{"success": "Username has been changed"})
		return
	}

	c.Redirect(http.StatusFound, "/admin")
}

// Функция добавления роли менеджера админом
func (db *DataBase) AddRole(c *gin.Context) {
	var role models.RoleAction
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	queryInsertUsersRole := `INSERT INTO users_roles (user_id, role_id) SELECT users_data.user_id, roles.role_id FROM users_data, roles WHERE users_data.user_id=$1 and roles.role_name=$2;`
	db.Data.MustExec(queryInsertUsersRole, &role.User_id, &role.Role)
	c.Redirect(http.StatusFound, "/admin")
}

// Функция удаления роли менеджера админом
func (db *DataBase) DeleteRole(c *gin.Context) {
	var role models.RoleAction
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	queryInsertUsersRole := `DELETE FROM users_roles WHERE user_id=$1 and role_id=(SELECT role_id FROM roles WHERE role_name = $2)`
	db.Data.MustExec(queryInsertUsersRole, &role.User_id, &role.Role)
	c.Redirect(http.StatusFound, "/admin")
}

// Функция удаления пользователя админом
func (db *DataBase) DeleteUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	queryInsertUsersRole := `DELETE FROM users_data WHERE user_id=$1`
	db.Data.MustExec(queryInsertUsersRole, &user.Id)
	c.Redirect(http.StatusFound, "/admin")
}
