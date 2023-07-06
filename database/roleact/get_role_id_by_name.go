package roleact

import (
	"github.com/jmoiron/sqlx"
)

func GetRoleIdByName(db *sqlx.DB, role_name string) (int, error) {
	var role_id int
	getrole := "select role_name from roles where role_name=$1"
	row, err := db.Query(getrole, role_name)
	if err != nil {
		return 0, err
	}

	defer row.Close()
	for row.Next() {
		if err := row.Scan(&role_id); err != nil {
			return 0, err
		}
	}
	return role_id, nil
}
