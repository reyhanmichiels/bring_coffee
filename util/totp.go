package util

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/pquerna/otp/totp"
	"gopkg.in/gomail.v2"
)

var validateOptions = totp.ValidateOpts{
	Period: 60 * 5,
	Digits: 4,
}

func GenerateOTP() (string, error) {
	code, err := totp.GenerateCodeCustom(os.Getenv("OTP_SECRET_KEY"), time.Now(), validateOptions)
	if err != nil {
		return "", err
	}

	return code, nil
}

func ValidateOTP(code string) (bool, error) {
	ok, err := totp.ValidateCustom(code, os.Getenv("OTP_SECRET_KEY"), time.Now(), validateOptions)
	if err != nil {
		return false, err
	}

	return ok, nil
}

func SendOTP(name, email, code string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "Bring Coffee")
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
