package access

import (
	"AdminPanelCorp/utils"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func GetAccessesToRole(db *sqlx.DB, roles []int, target_roles []int) bool {
	var rl int
	var all_roles []int
	getuser_accesses := "select access_to from access_roles where role_id = any ($1);"
	roww, err := db.Query(getuser_accesses, pq.Array(roles))
	if err != nil {
		return false
	}
	defer roww.Close()
	for roww.Next() {
		if err := roww.Scan(&rl); err != nil {
			return false
		}
		all_roles = append(all_roles, rl)
	}

	check_list := utils.Intersection(all_roles, target_roles)
	if len(check_list) == len(target_roles) {
		return true
	}

	return false
}
