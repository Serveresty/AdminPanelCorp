package utils

import (
	"AdminPanelCorp/env"
	"net/smtp"
)

// Функция отправки данных пользователю на почту
func SendEmail(data [][]string) [][]string {
	var emailErrors [][]string
	from := env.GetEnv("MAIL")
	pass := env.GetEnv("MAIL_PASSWORD")
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	auth := smtp.PlainAuth("", from, pass, host)
	var msg []byte
	var m []string
	for i := 0; i < len(data); i++ {
		msg = []byte("From: " + from + "\r\n" +
			"To: " + data[i][0] + "\r\n" + //
			"Subject: Your access data\r\n\r\n" +
			"Username: " + data[i][1] + "\n" + "Password: " + data[i][2] + "\r\n")

		m = append(m, data[i][0])
		err := smtp.SendMail(address, auth, from, m, msg)
		if err != nil {
			emailErrors = append(emailErrors, []string{data[i][0], "not valid email"})
		}
		m = nil
	}
	return emailErrors
}
