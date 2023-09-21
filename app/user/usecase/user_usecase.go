package usecase

import (
	"errors"
	"net/http"
	"strings"

	user_repository "github.com/reyhanmichiels/bring_coffee/app/user/repository"
	"github.com/reyhanmichiels/bring_coffee/domain"
	"github.com/reyhanmichiels/bring_coffee/util"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	RegistrationUsecase(request domain.RegistBind) (string, interface{})
	VerifyAccountUsecase(request domain.VerifyAccountBind) interface{}
	SendOTPUsecase(request domain.SendOTPBind) interface{}
}

type UserUsecase struct {
	UserRepo user_repository.IUserRepository
}

func NewUserUsecase(userRepo user_repository.IUserRepository) IUserUsecase {
	return &UserUsecase{
		UserRepo: userRepo,
	}
}

func (userUsecase *UserUsecase) RegistrationUsecase(request domain.RegistBind) (string, interface{}) {
	if request.Password != request.Verification_Password {
		return "", util.ErrorObject{
			Code:    http.StatusBadRequest,
			Message: "password and verification password doesn't same",
			Err:     errors.New(""),
		}
	}

	id, err := util.GenerateUlid()
	if err != nil {
		return "", util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to generate ulid",
			Err:     errors.New(""),
		}
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return "", util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to hash password",
			Err:     err,
		}
	}

	user := domain.User{
		ID:          id.String(),
		Name:        request.Name,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Password:    string(password),
	}

	err = userUsecase.UserRepo.CreateUser(&user)
	if err != nil {
		code := http.StatusInternalServerError
		if strings.Contains(err.Error(), "Duplicate entry") {
			code = http.StatusBadRequest
		}

		return "", util.ErrorObject{
			Code:    code,
			Message: "failed to create user",
			Err:     err,
		}
	}

	code, err := util.GenerateOTP()
	if err != nil {
		return "", util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to generate OTP",
			Err:     err,
		}
	}

	err = util.SendOTP(user.Name, user.Email, code)
	if err != nil {
		return "", util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to send OTP",
			Err:     err,
		}
	}

	return user.ID, nil
}

func (userUsecase *UserUsecase) VerifyAccountUsecase(request domain.VerifyAccountBind) interface{} {
	ok, err := util.ValidateOTP(request.Code)
	if err != nil {
		return util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "otp validation failed",
			Err:     err,
		}
	}

	if !ok {
		return util.ErrorObject{
			Code:    http.StatusBadRequest,
			Message: "otp doesn't same",
			Err:     errors.New(""),
		}
	}

	err = userUsecase.UserRepo.ActivateAccount(request.Email)
	if err != nil {
		return util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed activate account",
			Err:     err,
		}
	}

	return nil
}

func (userUsecase *UserUsecase) SendOTPUsecase(request domain.SendOTPBind) interface{} {
	user := struct {
		Name  string
		Email string
	}{}
	err := userUsecase.UserRepo.FindUserByCondition(&user, "email = ?", request.Email)
	if err != nil {
		return util.ErrorObject{
			Code:    http.StatusBadRequest,
			Message: "email not found",
			Err:     err,
		}
	}

	code, err := util.GenerateOTP()
	if err != nil {
		return util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to generate OTP",
			Err:     err,
		}
	}

	err = util.SendOTP(user.Name, user.Email, code)
	if err != nil {
		return util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to send OTP",
			Err:     err,
		}
	}

	return nil
}
