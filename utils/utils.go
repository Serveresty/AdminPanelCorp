package utils

import (
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) //Хэширование пароля
	return string(bytes), err
}

func Send_Email(data [][]string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	from := os.Getenv("MAIL")
	pass := os.Getenv("MAIL_PASSWORD")
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
			fmt.Println(err)
			os.Exit(1)
		}
		m = nil
	}
}

func Generate_Password() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 10
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
