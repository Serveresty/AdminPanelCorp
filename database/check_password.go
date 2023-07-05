package database

import (
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

// Функция проверки соответствия введенного пароля и пароля в БД
func CheckPassword(db *sqlx.DB, email string, password string) error {
	var hash_pass string
	getusers_hash_pass := "select password from users_data where email = $1"
	row, err := db.Query(getusers_hash_pass, email)
	if err != nil {
		return err
	}
	defer row.Close()
	for row.Next() {
		if err := row.Scan(&hash_pass); err != nil {
			return err
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash_pass), []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return err
		}
		return err
	}
	return nil
}
