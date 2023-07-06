package useract

import (
	"AdminPanelCorp/models"

	"github.com/jmoiron/sqlx"
)

func SetUsername(db *sqlx.DB, user models.User) {
	queryInsertNewUsersData := `UPDATE users_data SET username = $1 WHERE user_id=$2`
	db.MustExec(queryInsertNewUsersData, &user.Username, &user.Id)
}
