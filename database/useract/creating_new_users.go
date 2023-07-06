package useract

import (
	"AdminPanelCorp/utils"

	"github.com/jmoiron/sqlx"
)

func CreateUsers(db *sqlx.DB, records [][]string) ([][]string, [][]string) {
	var id int
	var result [][]string
	var email_errors [][]string
	queryInsertNewUser := `INSERT INTO "users_data" (email, username, password) VALUES($1, $2, $3)`
	getuser := "select user_id from users_data where username = $1"
	queryInsertUsersRole := `INSERT INTO "users_roles" (user_id, role_id) VALUES($1, $2)`

	for _, element := range records {
		if utils.IsEmailValid(element[0]) {
			password := utils.GeneratePassword()                                   //Генерация пароля
			hash_password, _ := utils.HashPassword(password)                       //Хэш пароля
			db.MustExec(queryInsertNewUser, element[0], element[1], hash_password) //Добавление в базу нового пользователя
			db.Get(&id, getuser, element[1])                                       //Получить id пользователя по username
			db.MustExec(queryInsertUsersRole, &id, 1)                              //Присвоение роли user по id пользователя
			result = append(result, []string{element[0], element[1], password})
		} else {
			email_errors = append(email_errors, []string{element[0], "not valid email"})
		}
	}
	return result, email_errors //Возвращаем данные вида (email, username, password) для отправки на почты пользователей
}
