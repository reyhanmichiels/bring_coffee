package usecase

import (
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
	return nil
}
