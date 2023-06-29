package server

import "github.com/jmoiron/sqlx"

// Структура с указателем на БД
type DataBase struct {
	Data *sqlx.DB
}
