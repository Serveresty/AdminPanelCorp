package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Функция создания таблиц с пользователями, ролями и их соответствием
func CreateTable(db *sqlx.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS "users_data" (user_id bigserial primary key, email varchar(255) unique, username varchar(255) unique, password varchar(255)); 
	CREATE TABLE IF NOT EXISTS "roles" (role_id serial primary key, role_name varchar(255) unique); 
	CREATE TABLE IF NOT EXISTS "users_roles" (user_id bigint references users_data (user_id) on delete cascade, role_id int references roles (role_id) on delete cascade);
	CREATE TABLE IF NOT EXISTS "access_roles" (role_id int references roles (role_id) on delete cascade, access_to int references roles (role_id) on delete cascade)`)
	return err
}
