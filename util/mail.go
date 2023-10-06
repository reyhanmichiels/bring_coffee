package util

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type IMail interface {
	SendOTP(name, email, code string) error
}

type Mail struct {
}

func NewMail() IMail {
	return &Mail{}
}

func (mail *Mail) SendOTP(name, email, code string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("SMTP_EMAIL"))
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "Your OTP Code")
	body := fmt.Sprintf("Hello %s, Welcome to Bring Coffee\n This is your OTP Code\n%s\nThis code only valid for 5 minutes", name, code)
	mailer.SetBody("text/html", body)

	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return err
	}

	dialer := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		smtpPort,
		os.Getenv("SMTP_EMAIL"),
		os.Getenv("SMTP_PASSWORD"),
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}
