package roleact

import (
	"github.com/jmoiron/sqlx"
)

func GetRoleIdByName(db *sqlx.DB, roleName string) (int, error) {
	var roleId int
	getrole := "select role_name from roles where role_name=$1"
	row, err := db.Query(getrole, roleName)
	if err != nil {
		return 0, err
	}

	defer row.Close()
	for row.Next() {
		if err := row.Scan(&roleId); err != nil {
			return 0, err
		}
	}
	return roleId, nil
}
