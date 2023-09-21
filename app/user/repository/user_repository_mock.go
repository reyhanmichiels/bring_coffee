package repository

import (
	"github.com/reyhanmichiels/bring_coffee/domain"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (userRepo *UserRepositoryMock) CreateUser(user *domain.User) error {
	return nil
}

func (userRepo *UserRepositoryMock) ActivateAccount(userEmail string) error {
	args := userRepo.Mock.Called(userEmail)
	if args[0] != nil {
		return args[0].(error)
	}

	return nil
}

func (userRepo *UserRepositoryMock) FindUserByCondition(user interface{}, conditon string, value interface{}) error {
	return nil
}
