package util

import (
	"os"
	"time"

	"github.com/pquerna/otp/totp"
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
