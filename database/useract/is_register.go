package useract

import "github.com/jmoiron/sqlx"

// Функция проверки на существующего пользователя
func IsUserRegistered(db *sqlx.DB, email string, username string) bool {
	var userID int
	getUserId := "select user_id from users_data where email = $1 or username = $2"
	row, err := db.Query(getUserId, email, username)
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
