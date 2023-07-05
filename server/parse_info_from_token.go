package server

import (
	"AdminPanelCorp/models"
	"AdminPanelCorp/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func parseInfoFromToken(c *gin.Context) (*models.Claims, string) {
	token := c.GetHeader("Authorization")
	if token == "" {
		return nil, "unauthorized"
	}
	token_string := strings.Split(token, " ")[1]

	claims, err2 := utils.ParseToken(token_string)

	if err2 != nil {
		return nil, "unauthorized"
	}
	return claims, ""
}
