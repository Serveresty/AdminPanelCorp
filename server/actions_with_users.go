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

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}
	token_string := strings.Split(token, " ")[1]

	claims, err2 := utils.ParseToken(token_string)

	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}

	auth_user, e := utils.GetUserByEmail(db.Data, claims.StandardClaims.Subject)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error db while get user"})
		return
	}

	target, err := utils.GetUsersRoles(db.Data, user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error db while get roles"})
		return
	}

	access := utils.CheckAccess(auth_user.Role, target)
	if access {
		if strings.ReplaceAll(user.Email, " ", "") != "" && strings.ReplaceAll(user.Username, " ", "") != "" {
			if utils.IsEmailValid(user.Email) {
				utils.SetEmail(db.Data, user)
				utils.SetUsername(db.Data, user)
				c.JSON(http.StatusOK, gin.H{"success": "Username and Email has been changed"})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "not valid email"})
			}
			return
		}
		if strings.ReplaceAll(user.Email, " ", "") != "" && strings.ReplaceAll(user.Username, " ", "") == "" {
			if utils.IsEmailValid(user.Email) {
				utils.SetEmail(db.Data, user)
				c.JSON(http.StatusOK, gin.H{"success": "Email has been changed"})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "not valid email"})
			}
			return
		}
		if strings.ReplaceAll(user.Email, " ", "") == "" && strings.ReplaceAll(user.Username, " ", "") != "" {
			utils.SetUsername(db.Data, user)
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

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}
	token_string := strings.Split(token, " ")[1]

	claims, err2 := utils.ParseToken(token_string)

	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}

	auth_user, e := utils.GetUserByEmail(db.Data, claims.StandardClaims.Subject)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error db while get user"})
		return
	}

	var access bool
	for _, elem := range auth_user.Role {
		if elem == "admin" {
			access = true
		}
	}
	if !access {
		c.JSON(http.StatusForbidden, gin.H{"error": "No rights to add role"})
		return
	}

	utils.AddRoleToUser(db.Data, role)
	c.JSON(http.StatusOK, gin.H{"success": "role has been added"})
}

// Функция удаления роли менеджера админом
func (db *DataBase) DeleteRole(c *gin.Context) {
	var role models.RoleAction
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}
	token_string := strings.Split(token, " ")[1]

	claims, err2 := utils.ParseToken(token_string)

	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}

	auth_user, e := utils.GetUserByEmail(db.Data, claims.StandardClaims.Subject)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error db while get user"})
		return
	}

	var access bool
	for _, elem := range auth_user.Role {
		if elem == "admin" {
			access = true
		}
	}
	if !access {
		c.JSON(http.StatusForbidden, gin.H{"error": "No rights to add role"})
		return
	}

	utils.DeleteRoleFromUser(db.Data, role)
	c.JSON(http.StatusOK, gin.H{"success": "role has been added"})
}

// Функция удаления пользователя админом
func (db *DataBase) DeleteUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}
	token_string := strings.Split(token, " ")[1]

	claims, err2 := utils.ParseToken(token_string)

	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}

	auth_user, e := utils.GetUserByEmail(db.Data, claims.StandardClaims.Subject)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error db while get user"})
		return
	}

	target, err := utils.GetUsersRoles(db.Data, user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error db while get roles"})
		return
	}

	access := utils.CheckAccess(auth_user.Role, target)

	if !access {
		c.JSON(http.StatusForbidden, gin.H{"error": "no rights to delete this user"})
		return
	}

	utils.DeleteUser(db.Data, user)
	c.JSON(http.StatusOK, gin.H{"success": "user has been deleted"})
}
