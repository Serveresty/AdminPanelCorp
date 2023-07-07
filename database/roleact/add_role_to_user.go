package roleact

import (
	"AdminPanelCorp/models"

	"github.com/jmoiron/sqlx"
)

func AddRoleToUser(db *sqlx.DB, role models.RoleAction) {
	queryInsertUsersRole := `INSERT INTO users_roles (user_id, role_id) SELECT users_data.user_id, roles.role_id FROM users_data, roles WHERE users_data.user_id=$1 and roles.role_name=$2;`
	db.MustExec(queryInsertUsersRole, &role.UserId, &role.Role)
}
