package utils

import (
	"AdminPanelCorp/models"

	"github.com/jmoiron/sqlx"
)

func DeleteUser(db *sqlx.DB, user models.User) {
	queryInsertUsersRole := `DELETE FROM users_data WHERE user_id=$1`
	db.MustExec(queryInsertUsersRole, &user.Id)
}
