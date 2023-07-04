package utils

import (
	"github.com/jmoiron/sqlx"
)

func GetUsersRoles(db *sqlx.DB, id int) ([]string, error) {
	var role string
	var all_roles []string
	getuser_roles := "select roles.role_name from users_data join users_roles on (users_roles.user_id=users_data.user_id) join roles on (roles.role_id=users_roles.role_id) where users_data.user_id=$1;"
	roww, err := db.Query(getuser_roles, id)
	if err != nil {
		return nil, err
	}
	defer roww.Close()
	for roww.Next() {
		if err := roww.Scan(&role); err != nil {
			return nil, err
		}
		all_roles = append(all_roles, role)
		role = ""
	}
	return all_roles, nil
}
