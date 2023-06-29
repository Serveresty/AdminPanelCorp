package database

import (
	"AdminPanelCorp/models"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// Функция инициализации БД
func DB_Init(cfg *models.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Функция создания таблиц с пользователями, ролями и их соответствием
func Create_Table(db *sqlx.DB) (error, error, error) {
	_, err1 := db.Exec(`CREATE TABLE IF NOT EXISTS "users_data" (user_id bigserial primary key, email varchar(255) unique, username varchar(255) unique, password varchar(255))`)
	_, err2 := db.Exec(`CREATE TABLE IF NOT EXISTS "roles" (role_id bigserial primary key, role_name varchar(255) unique)`)
	_, err3 := db.Exec(`CREATE TABLE IF NOT EXISTS "users_roles" (user_id bigint references users_data (user_id) on delete cascade, role_id bigint)`)
	return err1, err2, err3
}

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
