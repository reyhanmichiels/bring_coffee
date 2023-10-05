package usecase

import (
	"fmt"
	"testing"

	"github.com/reyhanmichiels/bring_coffee/app/mail"
	"github.com/reyhanmichiels/bring_coffee/app/user/repository"
	"github.com/reyhanmichiels/bring_coffee/domain"
	"github.com/reyhanmichiels/bring_coffee/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepositoryMock = repository.UserRepositoryMock{
	Mock: mock.Mock{},
}

var mailMock = mail.MailMock{
	Mock: mock.Mock{},
}

var userUsecase = NewUserUsecase(&userRepositoryMock, &mailMock)

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

func TestRegistrationUsecaseSuccessInput(t *testing.T) {
	request := []domain.RegistBind{
		{
			Name:                  "test",
			PhoneNumber:           "000000000000",
			Email:                 "test@test.com",
			Password:              "testpass",
			Verification_Password: "testpass",
		},
		{
			Name:                  "test",
			PhoneNumber:           "000000000000",
			Email:                 "test@test.com",
			Password:              "testpass",
			Verification_Password: "testpass",
		},
		{
			Name:                  "test",
			PhoneNumber:           "000000000000",
			Email:                 "test@test.com",
			Password:              "testpass",
			Verification_Password: "testpass",
		},
		{
			Name:                  "test",
			PhoneNumber:           "000000000000",
			Email:                 "test@test.com",
			Password:              "testpass",
			Verification_Password: "testpass",
		},
		{
			Name:                  "test",
			PhoneNumber:           "000000000000",
			Email:                 "test@test.com",
			Password:              "testpass",
			Verification_Password: "testpass",
		},
	}

	for i, v := range request {
		t.Run(fmt.Sprintf("feat: registration (usecase), test: Success Input %d", i+1), func(t *testing.T) {
			functionCall := userRepositoryMock.Mock.On("CreateUser", mock.Anything).Return(nil)
			functionCall2 := mailMock.Mock.On("SendOTP", v.Name, v.Email, mock.Anything).Return(nil)

			result, errObject := userUsecase.RegistrationUsecase(v)

			assert.Nil(t, errObject)
			assert.NotNil(t, result)

			functionCall.Unset()
			functionCall2.Unset()
		})
	}

}
