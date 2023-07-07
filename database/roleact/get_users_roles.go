package roleact

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func GetUsersRoles(db *sqlx.DB, userId int) ([]string, error) {
	var roleName string
	var roleNameArr []string
	roleIdArr, err := GetIdUsersRoles(db, userId)
	if err != nil {
		return nil, err
	}

	getRoleName := `select role_name from roles where role_id = any ($1)`
	roww, err := db.Query(getRoleName, pq.Array(roleIdArr))
	if err != nil {
		return nil, err
	}

	for roww.Next() {
		if err := roww.Scan(&roleName); err != nil {
			return nil, err
		}
		roleNameArr = append(roleNameArr, roleName)
	}
	return roleNameArr, nil
}

func GetIdUsersRoles(db *sqlx.DB, userId int) ([]int, error) {
	var roleId int
	var roleIdArr []int
	getRolesId := `SELECT role_id FROM users_roles WHERE user_id=$1`
	roww, err := db.Query(getRolesId, userId)
	if err != nil {
		return nil, err
	}
	defer roww.Close()
	for roww.Next() {
		if err := roww.Scan(&roleId); err != nil {
			return nil, err
		}
		roleIdArr = append(roleIdArr, roleId)
	}
	return roleIdArr, nil
}
