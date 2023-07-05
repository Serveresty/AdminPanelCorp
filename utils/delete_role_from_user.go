package utils

import (
	"AdminPanelCorp/models"

	"github.com/jmoiron/sqlx"
)

func DeleteRoleFromUser(db *sqlx.DB, role models.RoleAction) {
	queryInsertUsersRole := `DELETE FROM users_roles WHERE user_id=$1 and role_id=(SELECT role_id FROM roles WHERE role_name = $2)`
	db.MustExec(queryInsertUsersRole, &role.User_id, &role.Role)
}
