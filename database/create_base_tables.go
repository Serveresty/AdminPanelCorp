package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Функция создания таблиц с пользователями, ролями и их соответствием
func CreateTable(db *sqlx.DB) (error, error, error) {
	_, err1 := db.Exec(`CREATE TABLE IF NOT EXISTS "users_data" (user_id bigserial primary key, email varchar(255) unique, username varchar(255) unique, password varchar(255))`)
	_, err2 := db.Exec(`CREATE TABLE IF NOT EXISTS "roles" (role_id bigserial primary key, role_name varchar(255) unique)`)
	_, err3 := db.Exec(`CREATE TABLE IF NOT EXISTS "users_roles" (user_id bigint references users_data (user_id) on delete cascade, role_id bigint)`)
	return err1, err2, err3
}
