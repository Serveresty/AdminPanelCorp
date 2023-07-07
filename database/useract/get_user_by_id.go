package useract

import (
	"AdminPanelCorp/database/roleact"
	"AdminPanelCorp/models"

	"github.com/jmoiron/sqlx"
)

func GetUserById(db *sqlx.DB, id int) (models.User, error) {
	var usersData models.User
	getuser := "select user_id, email, username, password from users_data where user_id=$1"
	row, err := db.Query(getuser, id)
	if err != nil {
		return models.User{}, err
	}

	defer row.Close()
	for row.Next() {
		if err := row.Scan(&usersData.Id, &usersData.Email, &usersData.Username, &usersData.Password); err != nil {
			return models.User{}, err
		}

		roles, errRoles := roleact.GetUsersRoles(db, usersData.Id)
		if errRoles != nil {
			return models.User{}, errRoles
		}
		usersData.Role = roles

	}
	return usersData, nil
}
