package access

import (
	"AdminPanelCorp/database/roleact"

	"github.com/jmoiron/sqlx"
)

func AddAccessesToRole(db *sqlx.DB, role string, access_role []string) error {
	for _, elem := range access_role {
		if !isAccessAlreadyGranted(db, role, elem) {
			role_id, err1 := roleact.GetRoleIdByName(db, role)
			if err1 != nil {
				return err1
			}
			elem_id, err2 := roleact.GetRoleIdByName(db, elem)
			if err2 != nil {
				return err2
			}
			queryInsertNewAccess := `INSERT INTO "access_roles" (role_id, access_to) VALUES($1, $2)`
			db.MustExec(queryInsertNewAccess, role_id, elem_id)
		}
	}
	return nil
}

func isAccessAlreadyGranted(db *sqlx.DB, role string, target_role string) bool {
	var flag int
	get_user_id := "select role_id from access_roles where role_id = $1 and access_to = $2"
	row, err := db.Query(get_user_id, role, target_role)
	if err != nil {
		return false
	}
	defer row.Close()
	for row.Next() {
		if err := row.Scan(&flag); err != nil {
			return false
		}
	}
	if flag != 0 {
		return true
	}
	return false
}
