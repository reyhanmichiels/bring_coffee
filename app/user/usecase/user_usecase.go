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
	BasicLoginUsecase(request domain.BasicLoginBind) (interface{}, interface{})
	VerifyForgetPasswordUsecase(request domain.VerifyAccountBind) (interface{}, interface{})
	ForgetPasswordUsecase(email string, request domain.ForgetPasswordBind) interface{}
}

type UserUsecase struct {
	UserRepo user_repository.IUserRepository
	Mail     util.IMail
}

func NewUserUsecase(userRepo user_repository.IUserRepository, mail util.IMail) IUserUsecase {
	return &UserUsecase{
		UserRepo: userRepo,
		Mail:     mail,
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

	err = userUsecase.Mail.SendOTP(user.Name, user.Email, code)
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

	err = userUsecase.Mail.SendOTP(user.Name, user.Email, code)
	if err != nil {
		return util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to send OTP",
			Err:     err,
		}
	}

	return nil
}

func (userUsecase *UserUsecase) BasicLoginUsecase(request domain.BasicLoginBind) (interface{}, interface{}) {
	var user domain.User
	err := userUsecase.UserRepo.FindUserByCondition(&user, "email = ?", request.Email)
	if err != nil {
		return nil, util.ErrorObject{
			Code:    http.StatusBadRequest,
			Message: "account not found",
			Err:     err,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, util.ErrorObject{
			Code:    http.StatusBadRequest,
			Message: "credential doesn't match",
			Err:     err,
		}
	}

	if !user.IsVerified {
		return nil, util.ErrorObject{
			Code:    http.StatusUnauthorized,
			Message: "user is not verified",
			Err:     errors.New(""),
		}
	}

	token, err := util.GenerateJWT(user.ID)
	if err != nil {
		return nil, util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to generate jwt token",
			Err:     err,
		}
	}

	apiResponse := struct {
		Token string `json:"token"`
	}{
		token,
	}
	return apiResponse, nil
}

func (userUsecase *UserUsecase) VerifyForgetPasswordUsecase(request domain.VerifyAccountBind) (interface{}, interface{}) {
	ok, err := util.ValidateOTP(request.Code)
	if err != nil {
		return nil, util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "otp validation failed",
			Err:     err,
		}
	}

	if !ok {
		return nil, util.ErrorObject{
			Code:    http.StatusBadRequest,
			Message: "otp doesn't same",
			Err:     errors.New(""),
		}
	}

	token, err := util.GenerateTokenForgetPassword(request.Email)
	if err != nil {
		return nil, util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to generate OTP",
			Err:     err,
		}
	}

	apiResponse := struct {
		Token string `json:"token"`
	}{
		token,
	}
	return apiResponse, nil
}

func (userUsecase *UserUsecase) ForgetPasswordUsecase(email string, request domain.ForgetPasswordBind) interface{} {
	if request.Password != request.Verification_Password {
		return util.ErrorObject{
			Code:    http.StatusBadRequest,
			Message: "password and verification password doesn't same",
			Err:     errors.New(""),
		}
	}

	var user domain.User
	err := userUsecase.UserRepo.FindUserByCondition(&user, "email = ?", email)
	if err != nil {
		return util.ErrorObject{
			Code:    http.StatusBadRequest,
			Message: "email not found",
			Err:     err,
		}
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to hash password",
			Err:     err,
		}
	}

	userUpdateData := struct {
		Password string
	}{
		string(password),
	}

	err = userUsecase.UserRepo.Update(&user, userUpdateData)
	if err != nil {
		return util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to change password",
			Err:     err,
		}
	}

	return nil
}
