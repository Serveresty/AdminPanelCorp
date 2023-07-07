package access

import (
	"AdminPanelCorp/database/roleact"

	"github.com/jmoiron/sqlx"
)

func AddAccessesToRole(db *sqlx.DB, role string, access_role []string) error {
	for _, elem := range access_role {
		if !isAccessAlreadyGranted(db, role, elem) {
			roleId, err1 := roleact.GetRoleIdByName(db, role)
			if err1 != nil {
				return err1
			}
			elemId, err2 := roleact.GetRoleIdByName(db, elem)
			if err2 != nil {
				return err2
			}
			queryInsertNewAccess := `INSERT INTO "access_roles" (role_id, access_to) VALUES($1, $2)`
			db.MustExec(queryInsertNewAccess, roleId, elemId)
		}
	}
	return nil
}

func isAccessAlreadyGranted(db *sqlx.DB, role string, targetRole string) bool {
	var flag int
	getUserId := "select role_id from access_roles where role_id = $1 and access_to = $2"
	row, err := db.Query(getUserId, role, targetRole)
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
