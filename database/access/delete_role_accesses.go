package access

import (
	"AdminPanelCorp/database/roleact"

	"github.com/jmoiron/sqlx"
)

func DeleteAccessesFromRole(db *sqlx.DB, role string, accessRole []string) error {
	for _, elem := range accessRole {
		roleId, err1 := roleact.GetRoleIdByName(db, role)
		if err1 != nil {
			return err1
		}
		elemId, err2 := roleact.GetRoleIdByName(db, elem)
		if err2 != nil {
			return err2
		}
		if isAccessAlreadyGranted(db, roleId, elemId) {
			queryInsertNewAccess := `DELETE FROM access_roles WHERE role_id=$1 and access_to=$2`
			db.MustExec(queryInsertNewAccess, roleId, elemId)
		}
	}
	return nil
}
