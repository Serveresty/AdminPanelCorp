package useract

import (
	"AdminPanelCorp/utils"

	"github.com/jmoiron/sqlx"
)

func CreateUsers(db *sqlx.DB, records [][]string) ([][]string, [][]string) {
	var id int
	var result [][]string
	var emailErrors [][]string
	queryInsertNewUser := `INSERT INTO "users_data" (email, username, password) VALUES($1, $2, $3)`
	getuser := "select user_id from users_data where username = $1"
	queryInsertUsersRole := `INSERT INTO "users_roles" (user_id, role_id) VALUES($1, $2)`

	for _, element := range records {
		if utils.IsEmailValid(element[0]) {
			password := utils.GeneratePassword()                                  //Генерация пароля
			hashPassword, _ := utils.HashPassword(password)                       //Хэш пароля
			db.MustExec(queryInsertNewUser, element[0], element[1], hashPassword) //Добавление в базу нового пользователя
			db.Get(&id, getuser, element[1])                                      //Получить id пользователя по username
			db.MustExec(queryInsertUsersRole, &id, 1)                             //Присвоение роли user по id пользователя
			result = append(result, []string{element[0], element[1], password})
		} else {
			emailErrors = append(emailErrors, []string{element[0], "not valid email"})
		}
	}
	return result, emailErrors //Возвращаем данные вида (email, username, password) для отправки на почты пользователей
}
