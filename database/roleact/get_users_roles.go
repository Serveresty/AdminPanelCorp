package roleact

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func GetUsersRoles(db *sqlx.DB, user_id int) ([]string, error) {
	var role_name string
	var role_name_arr []string
	role_id_arr, err := GetIdUsersRoles(db, user_id)
	if err != nil {
		return nil, err
	}

	get_role_name := `select role_name from roles where role_id = any ($1)`
	roww, err := db.Query(get_role_name, pq.Array(role_id_arr))
	if err != nil {
		return nil, err
	}

	for roww.Next() {
		if err := roww.Scan(&role_name); err != nil {
			return nil, err
		}
		role_name_arr = append(role_name_arr, role_name)
	}
	return role_name_arr, nil
}

func GetIdUsersRoles(db *sqlx.DB, user_id int) ([]int, error) {
	var role_id int
	var role_id_arr []int
	get_roles_id := `SELECT role_id FROM users_roles WHERE user_id=$1`
	roww, err := db.Query(get_roles_id, user_id)
	if err != nil {
		return nil, err
	}
	defer roww.Close()
	for roww.Next() {
		if err := roww.Scan(&role_id); err != nil {
			return nil, err
		}
		role_id_arr = append(role_id_arr, role_id)
	}
	return role_id_arr, nil
}
