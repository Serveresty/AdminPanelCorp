package database

import (
	"AdminPanelCorp/models"

	"github.com/jmoiron/sqlx"
)

func GetUserByEmail(db *sqlx.DB, email string) (models.User, error) {
	var users_data models.User
	getuser := "select user_id, email, username, password from users_data where email=$1"
	row, err := db.Query(getuser, email)
	if err != nil {
		return models.User{}, err
	}

	defer row.Close()
	for row.Next() {
		if err := row.Scan(&users_data.Id, &users_data.Email, &users_data.Username, &users_data.Password); err != nil {
			return models.User{}, err
		}

		roles, err_roles := GetUsersRoles(db, users_data.Id)
		if err_roles != nil {
			return models.User{}, err_roles
		}
		users_data.Role = roles

	}
	return users_data, nil
}