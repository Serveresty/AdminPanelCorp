package useract

import (
	"AdminPanelCorp/database/roleact"
	"AdminPanelCorp/models"

	"github.com/jmoiron/sqlx"
)

// Функция для админ панели на получение всех зарегистрированных пользователей
func GetAllUsers(db *sqlx.DB) ([]models.User, error) {
	var usersData []models.User
	getuser := "select user_id, email, username, password from users_data"
	row, err := db.Query(getuser)
	if err != nil {
		return nil, err
	}

	defer row.Close()
	for row.Next() {
		var currentUser models.User
		if err := row.Scan(&currentUser.Id, &currentUser.Email, &currentUser.Username, &currentUser.Password); err != nil {
			return nil, err
		}
		roles, err_roles := roleact.GetUsersRoles(db, currentUser.Id)
		if err_roles != nil {
			return nil, err_roles
		}
		currentUser.Role = roles

		usersData = append(usersData, currentUser)
	}
	return usersData, nil
}
