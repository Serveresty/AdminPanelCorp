package access

import (
	"AdminPanelCorp/utils"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func GetAccessesToRole(db *sqlx.DB, roles []int, targetRoles []int) bool {
	var rl int
	var allRoles []int
	getuserAccesses := "select access_to from access_roles where role_id = any ($1);"
	roww, err := db.Query(getuserAccesses, pq.Array(roles))
	if err != nil {
		return false
	}
	defer roww.Close()
	for roww.Next() {
		if err := roww.Scan(&rl); err != nil {
			return false
		}
		allRoles = append(allRoles, rl)
	}

	checkList := utils.Intersection(allRoles, targetRoles)
	if len(checkList) == len(targetRoles) {
		return true
	}

	return false
}
