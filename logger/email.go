package logger

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/gomail.v2"
)

type EmailConfig struct {
	From     string
	Password string
	Level    string
}

const (
	gmailSMTP = "smtp.gmail.com"
	gmailPort = 587
)

var emailConfig *EmailConfig

func EmailError(messageString, service string) {
	if emailConfig == nil {
		return
	}

	m := gomail.NewMessage()

	m.SetHeader("From", emailConfig.From)
	m.SetHeader("To", emailConfig.From)
	m.SetHeader("Subject", fmt.Sprintf("%s - %s", emailConfig.Level, service))

	m.SetBody("text/plain", messageString)

	d := gomail.NewDialer(gmailSMTP, gmailPort, emailConfig.From, emailConfig.Password)

	//nolint
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
