package access

import (
	"AdminPanelCorp/database/roleact"

	"github.com/jmoiron/sqlx"
)

func DeleteAccessesFromRole(db *sqlx.DB, role string, access_role []string) error {
	for _, elem := range access_role {
		if isAccessAlreadyGranted(db, role, elem) {
			role_id, err1 := roleact.GetRoleIdByName(db, role)
			if err1 != nil {
				return err1
			}
			elem_id, err2 := roleact.GetRoleIdByName(db, elem)
			if err2 != nil {
				return err2
			}
			queryInsertNewAccess := `DELETE FROM access_roles WHERE role_id=$1 and access_to=$2`
			db.MustExec(queryInsertNewAccess, role_id, elem_id)
		}
	}
	return nil
}
