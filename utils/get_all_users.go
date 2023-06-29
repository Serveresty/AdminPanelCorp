package utils

import (
	"AdminPanelCorp/models"

	"github.com/jmoiron/sqlx"
)

// Функция для админ панели на получение всех зарегистрированных пользователей
func GetAllUsers(db *sqlx.DB) ([]models.User, error) {
	var users_data []models.User
	getuser := "select users_data.user_id, email, username, password, roles.role_name from users_data join users_roles on (users_roles.user_id=users_data.user_id) join roles on (roles.role_id=users_roles.role_id)"
	row, err := db.Query(getuser)
	if err != nil {
		return nil, err
	}

	defer row.Close()
	for row.Next() {
		var current_user models.User
		if err := row.Scan(&current_user.Id, &current_user.Email, &current_user.Username, &current_user.Password, &current_user.Role); err != nil {
			return nil, err
		}
		users_data = append(users_data, current_user)
	}
	return users_data, nil
}
