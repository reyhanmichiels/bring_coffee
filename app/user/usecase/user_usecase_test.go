package usecase

import (
	"fmt"
	"testing"

	"github.com/reyhanmichiels/bring_coffee/app/user/repository"
	"github.com/reyhanmichiels/bring_coffee/domain"
	"github.com/reyhanmichiels/bring_coffee/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepositoryMock = repository.UserRepositoryMock{
	Mock: mock.Mock{},
}

var userUsecase = NewUserUsecase(&userRepositoryMock)

func TestVerifyAccountUsecaseSuccessInput(t *testing.T) {
	code, _ := util.GenerateOTP()

	request := []domain.VerifyAccountBind{
		{
			Email: "test1@test.com",
			Code:  code,
		},
		{
			Email: "test2@test.com",
			Code:  code,
		},
		{
			Email: "test3@test.com",
			Code:  code,
		},
		{
			Email: "test4@test.com",
			Code:  code,
		},
		{
			Email: "test5@test.com",
			Code:  code,
		},
	}

	for i, v := range request {
		t.Run(fmt.Sprintf("feat: verify account (usecase), test: Success Input %d", i+1), func(t *testing.T) {
			callFunction := userRepositoryMock.Mock.On("ActivateAccount", v.Email).Return(nil)

			successResponse := userUsecase.VerifyAccountUsecase(v)
			assert.Nil(t, successResponse, "response should be nil")

			callFunction.Unset()
		})
	}
}
