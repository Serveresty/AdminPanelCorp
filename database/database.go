package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func DB_Init(cfg Config) (*sqlx.DB, error) {
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

func Create_Table(db *sqlx.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS "users_data" (user_id bigserial primary key, email varchar(255) unique, username varchar(255) unique, password varchar(255))`)
	if err != nil {
		panic(err)
	}
	_, err1 := db.Exec(`CREATE TABLE IF NOT EXISTS "roles" (role_id bigserial primary key, role_name varchar(255) unique)`)
	if err1 != nil {
		panic(err1)
	}
	_, err2 := db.Exec(`CREATE TABLE IF NOT EXISTS "users_roles" (user_id bigint, role_id bigint)`)
	if err2 != nil {
		panic(err1)
	}
}

func IsUserRegistered(db *sqlx.DB, email string) bool {
	var userID int
	get_user := "select password from users_data where email = $1"
	db.Get(&userID, get_user, email)
	if userID != 0 {
		return true
	}
	return false
}

func CheckPassword(db *sqlx.DB, email string, password string) error {
	var hash_pass string
	getusers_hash_pass := "select password from users_data where email = $1"
	db.Get(&hash_pass, getusers_hash_pass, email)
	if err := bcrypt.CompareHashAndPassword([]byte(hash_pass), []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return err
		}
		return err
	}
	return nil
}
