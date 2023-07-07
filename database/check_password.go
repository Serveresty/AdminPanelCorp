package database

import (
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

// Функция проверки соответствия введенного пароля и пароля в БД
func CheckPassword(db *sqlx.DB, email string, password string) error {
	var hashPass string
	getusersHashPass := "select password from users_data where email = $1"
	row, err := db.Query(getusersHashPass, email)
	if err != nil {
		return err
	}
	defer row.Close()
	for row.Next() {
		if err := row.Scan(&hashPass); err != nil {
			return err
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return err
		}
		return err
	}
	return nil
}
