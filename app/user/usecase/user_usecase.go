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

	return user.ID, nil
}
