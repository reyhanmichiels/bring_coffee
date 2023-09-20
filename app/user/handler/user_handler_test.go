package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/reyhanmichiels/bring_coffee/app/user/usecase"
	"github.com/reyhanmichiels/bring_coffee/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userUsecaseMock = usecase.UserUsecaseMock{
	Mock: mock.Mock{},
}

var userHandler = NewUserHandler(&userUsecaseMock)

func TestRegistrationSuccessInput(t *testing.T) {
	request := []domain.RegistBind{
		{
			Name:                  "test1",
			PhoneNumber:           "1234567891234",
			Email:                 "test1@test.com",
			Password:              "testPassword1",
			Verification_Password: "testPassword1",
		},
		{
			Name:                  "test2",
			PhoneNumber:           "2234567892234",
			Email:                 "test2@test.com",
			Password:              "testPassword2",
			Verification_Password: "testPassword2",
		},
		{
			Name:                  "test3",
			PhoneNumber:           "3234567893234",
			Email:                 "test3@test.com",
			Password:              "testPassword3",
			Verification_Password: "testPassword3",
		},
		{
			Name:                  "test4",
			PhoneNumber:           "4234567894234",
			Email:                 "test4@test.com",
			Password:              "testPassword4",
			Verification_Password: "testPassword4",
		},
		{
			Name:                  "test5",
			PhoneNumber:           "5234567895234",
			Email:                 "test5@test.com",
			Password:              "testPassword5",
			Verification_Password: "testPassword5",
		},
	}

	for i, v := range request {
		t.Run(fmt.Sprintf("feat: registration (handler), test: Success Input %d", i+1), func(t *testing.T) {
			functionResponse := fmt.Sprint("test", 1)
			callFunction := userUsecaseMock.Mock.On("RegistrationUsecase", v).Return(functionResponse, nil)

			engine := gin.Default()
			engine.POST("/api/v1/regist", userHandler.Registration)

			requestDataInJson, err := json.Marshal(v)
			if err != nil {
				t.Fatal(err)
			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", "/api/v1/regist", bytes.NewBuffer(requestDataInJson))
			if err != nil {
				t.Fatal(err)
			}
			engine.ServeHTTP(response, request)

			var responseBody map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &responseBody)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, http.StatusCreated, response.Code, "http status code should be equal")
			assert.Equal(t, "success", responseBody["status"], "status should be equal")
			assert.Equal(t, "successfully registration new user", responseBody["message"], "message should be equal")
			assert.Equal(t, functionResponse, responseBody["data"], "data should be equal")

			callFunction.Unset()
		})
	}
}
