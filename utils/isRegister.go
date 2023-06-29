package utils

import "github.com/jmoiron/sqlx"

// Функция проверки на существующего пользователя
func IsUserRegistered(db *sqlx.DB, email string) bool {
	var userID int
	get_user_id := "select user_id from users_data where email = $1"
	row, err := db.Query(get_user_id, email)
	if err != nil {
		return false
	}
	defer row.Close()
	for row.Next() {
		if err := row.Scan(&userID); err != nil {
			return false
		}
	}
	if userID != 0 {
		return true
	}
	return false
}
