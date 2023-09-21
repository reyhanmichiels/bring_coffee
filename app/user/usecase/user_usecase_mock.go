package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/reyhanmichiels/bring_coffee/domain"
	"github.com/stretchr/testify/mock"
)

type UserUsecaseMock struct {
	Mock mock.Mock
}

func (userUsecase *UserUsecaseMock) RegistrationUsecase(request domain.RegistBind) (string, interface{}) {
	args := userUsecase.Mock.Called(request)
	if args[1] != nil {
		return "", args[1]
	}

	return args[0].(string), nil
}

func (userUsecase *UserUsecaseMock) VerifyAccountUsecase(request domain.VerifyAccountBind) interface{} {
	args := userUsecase.Mock.Called(request)
	if args[0] != nil {
		return args[0]
	}

	return nil
}

func (userUsecase *UserUsecaseMock) SendOTPUsecase(request domain.SendOTPBind) interface{} {
	args := userUsecase.Mock.Called(request)
	if args[0] != nil {
		return args[0]
	}

	return nil
}

func (userUsecase *UserUsecaseMock) BasicLoginUsecase(c *gin.Context, request domain.BasicLoginBind) (interface{}, interface{}) {
	args := userUsecase.Mock.Called(request)
	if args[1] != nil {
		return nil, args[1]
	}

	return args[0], nil
}
