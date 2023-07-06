package server

import (
	"AdminPanelCorp/database"
	"AdminPanelCorp/models"
	"AdminPanelCorp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func (db *DataBase) EditUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, err_str := parseInfoFromToken(c)
	if err_str != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}

	auth_user, e := database.GetUserByEmail(db.Data, claims.StandardClaims.Subject) //User from token
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error db while get user"})
		return
	}

	roles_id_auth, err := database.GetIdUsersRoles(db.Data, auth_user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error db while get roles"})
		return
	}

	target, err := database.GetIdUsersRoles(db.Data, user.Id) //The User being modified
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error db while get roles"})
		return
	}

	access := database.GetAccessesToRole(db.Data, roles_id_auth, target)
	if access {
		if user.Email != "" && user.Username != "" {
			if utils.IsEmailValid(user.Email) {
				database.SetEmail(db.Data, user)
				database.SetUsername(db.Data, user)
				c.JSON(http.StatusOK, gin.H{"success": "Username and Email has been changed"})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "not valid email"})
			}
			return
		}
		if user.Email != "" && user.Username == "" {
			if utils.IsEmailValid(user.Email) {
				database.SetEmail(db.Data, user)
				c.JSON(http.StatusOK, gin.H{"success": "Email has been changed"})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "not valid email"})
			}
			return
		}
		if user.Email == "" && user.Username != "" {
			database.SetUsername(db.Data, user)
			c.JSON(http.StatusOK, gin.H{"success": "Username has been changed"})
			return
		}
	}

	c.JSON(http.StatusForbidden, gin.H{"error": "no rights to edit this user"})
}

// Функция добавления роли менеджера админом
func (db *DataBase) AddRole(c *gin.Context) {
	var role models.RoleAction
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, err_str := parseInfoFromToken(c)
	if err_str != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}

	auth_user, e := database.GetUserByEmail(db.Data, claims.StandardClaims.Subject)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error db while get user"})
		return
	}

	roles_id_auth, err := database.GetIdUsersRoles(db.Data, auth_user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error db while get roles"})
		return
	}

	target, err := database.GetIdUsersRoles(db.Data, role.User_id) //The User being modified
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error db while get roles"})
		return
	}

	access := database.GetAccessesToRole(db.Data, roles_id_auth, target)

	if access {
		database.AddRoleToUser(db.Data, role)
		c.JSON(http.StatusOK, gin.H{"success": "role has been added"})
		return
	}

	c.JSON(http.StatusForbidden, gin.H{"error": "No rights to add role"})
}

// Функция удаления роли менеджера админом
func (db *DataBase) DeleteRole(c *gin.Context) {
	var role models.RoleAction
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, err_str := parseInfoFromToken(c)
	if err_str != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}

	auth_user, e := database.GetUserByEmail(db.Data, claims.StandardClaims.Subject)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error db while get user"})
		return
	}

	roles_id_auth, err := database.GetIdUsersRoles(db.Data, auth_user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error db while get roles"})
		return
	}

	target, err := database.GetIdUsersRoles(db.Data, role.User_id) //The User being modified
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error db while get roles"})
		return
	}

	access := database.GetAccessesToRole(db.Data, roles_id_auth, target)
	if access {
		database.DeleteRoleFromUser(db.Data, role)
		c.JSON(http.StatusOK, gin.H{"success": "role has been added"})
		return
	}

	c.JSON(http.StatusForbidden, gin.H{"error": "No rights to add role"})
}

// Функция удаления пользователя админом
func (db *DataBase) DeleteUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, err_str := parseInfoFromToken(c)
	if err_str != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}

	auth_user, e := database.GetUserByEmail(db.Data, claims.StandardClaims.Subject)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error db while get user"})
		return
	}

	roles_id_auth, err := database.GetIdUsersRoles(db.Data, auth_user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error db while get roles"})
		return
	}

	target, err := database.GetIdUsersRoles(db.Data, user.Id) //The User being modified
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error db while get roles"})
		return
	}

	access := database.GetAccessesToRole(db.Data, roles_id_auth, target)

	if access {
		database.DeleteUser(db.Data, user)
		c.JSON(http.StatusOK, gin.H{"success": "user has been deleted"})
		return
	}

	c.JSON(http.StatusForbidden, gin.H{"error": "no rights to delete this user"})
}
