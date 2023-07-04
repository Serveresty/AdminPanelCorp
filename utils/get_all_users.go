package utils

import (
	"AdminPanelCorp/models"

	"github.com/jmoiron/sqlx"
)

// Функция для админ панели на получение всех зарегистрированных пользователей
func GetAllUsers(db *sqlx.DB) ([]models.User, error) {
	var users_data []models.User
	getuser := "select user_id, email, username, password from users_data"
	row, err := db.Query(getuser)
	if err != nil {
		return nil, err
	}

	defer row.Close()
	for row.Next() {
		var current_user models.User
		if err := row.Scan(&current_user.Id, &current_user.Email, &current_user.Username, &current_user.Password); err != nil {
			return nil, err
		}

		roles, err_roles := GetUsersRoles(db, current_user.Email)
		if err_roles != nil {
			return nil, err_roles
		}
		current_user.Role = roles

		users_data = append(users_data, current_user)
	}
	return users_data, nil
}
