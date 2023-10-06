package repository

import (
	"github.com/reyhanmichiels/bring_coffee/domain"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (userRepositoryMock *UserRepositoryMock) CreateUser(user *domain.User) error {
	args := userRepositoryMock.Mock.Called(user)
	if args[0] != nil {
		return args[0].(error)
	}

	return nil
}

func (userRepositoryMock *UserRepositoryMock) ActivateAccount(userEmail string) error {
	args := userRepositoryMock.Mock.Called(userEmail)
	if args[0] != nil {
		return args[0].(error)
	}

	return nil
}

func (userRepositoryMock *UserRepositoryMock) FindUserByCondition(user interface{}, condition string, value interface{}) error {
	args := userRepositoryMock.Mock.Called(user, condition, value)
	if args[0] != nil {
		return args[0].(error)
	}

	user = &domain.User{
		Password: "password",
	}

	return nil
}

func (userRepositoryMock *UserRepositoryMock) Update(user *domain.User, userUpdateData interface{}) error {
	args := userRepositoryMock.Mock.Called(user, userUpdateData)
	if args[0] != nil {
		return args[0].(error)
	}

	user = &domain.User{
		Password: "password",
	}

	return nil
}