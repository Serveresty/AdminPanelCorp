package utils

import (
	"AdminPanelCorp/models"
	"fmt"

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
		var role string
		if err := row.Scan(&current_user.Id, &current_user.Email, &current_user.Username, &current_user.Password); err != nil {
			return nil, err
		}

		getuser_roles := "select roles.role_name from users_data join users_roles on (users_roles.user_id=users_data.user_id) join roles on (roles.role_id=users_roles.role_id) where users_data.email=$1;"
		roww, err := db.Query(getuser_roles, current_user.Email)
		if err != nil {
			fmt.Println("firsttt")
			fmt.Println(err)
		}
		defer roww.Close()
		for roww.Next() {
			if err := roww.Scan(&role); err != nil {
				fmt.Println("seconddd")
				fmt.Println(err)
			}
			current_user.Role = append(current_user.Role, role)
		}

		users_data = append(users_data, current_user)
		role = ""
	}
	return users_data, nil
}
