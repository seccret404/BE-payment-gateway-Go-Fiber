package utils

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-gomail/gomail"
)

func SendEmail(to, subject, body string) error {
	host := os.Getenv("SMTP_HOST")
	portStr := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	sender := os.Getenv("SMTP_SENDER_NAME") 

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return fmt.Errorf("invalid SMTP port: %v", err)
	}

	m := gomail.NewMessage()

	//fyi : Format harus "Nama <email@example.com>"
	m.SetHeader("From", m.FormatAddress(user, sender))
	m.SetHeader("To", to) // lowercase "To" !
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(host, port, user, pass)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed send email to %s: %v", to, err)
	}

	return nil
}
