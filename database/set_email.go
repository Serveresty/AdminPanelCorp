package database

import (
	"AdminPanelCorp/models"

	"github.com/jmoiron/sqlx"
)

func SetEmail(db *sqlx.DB, user models.User) {
	queryInsertNewUsersData := `UPDATE users_data SET email = $1 WHERE user_id=$2`
	db.MustExec(queryInsertNewUsersData, &user.Email, &user.Id)
}
