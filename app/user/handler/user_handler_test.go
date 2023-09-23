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
			engine.POST("/api/auth/regist", userHandler.Registration)

			requestDataInJson, err := json.Marshal(v)
			if err != nil {
				t.Fatal(err)
			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", "/api/auth/regist", bytes.NewBuffer(requestDataInJson))
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

func TestVerifyAccountSuccessInput(t *testing.T) {
	request := []domain.VerifyAccountBind{
		{
			Email: "test1@test.com",
			Code:  "1111",
		},
		{
			Email: "test2@test.com",
			Code:  "2111",
		},
		{
			Email: "test3@test.com",
			Code:  "3111",
		},
		{
			Email: "test4@test.com",
			Code:  "4111",
		},
		{
			Email: "test5@test.com",
			Code:  "5111",
		},
	}

	for i, v := range request {
		t.Run(fmt.Sprintf("feat: verify account (handler), test: Success Input %d", i+1), func(t *testing.T) {
			callFunction := userUsecaseMock.Mock.On("VerifyAccountUsecase", v).Return(nil)

			engine := gin.Default()
			engine.POST("/api/v1/users/verify", userHandler.VerifyAccount)

			requestDataInJson, err := json.Marshal(v)
			if err != nil {
				t.Fatal(err)
			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", "/api/v1/users/verify", bytes.NewBuffer(requestDataInJson))
			if err != nil {
				t.Fatal(err)
			}
			engine.ServeHTTP(response, request)

			var responseBody map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &responseBody)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, http.StatusOK, response.Code, "http status code should be equal")
			assert.Equal(t, "success", responseBody["status"], "status should be equal")
			assert.Equal(t, "successfully verified account", responseBody["message"], "message should be equal")
			assert.Nil(t, responseBody["data"], "data should be nil")

			callFunction.Unset()
		})
	}
}

func TestSendOTPSuccessInput(t *testing.T) {
	request := []domain.SendOTPBind{
		{
			Email: "test1@gmail.com",
		},
		{
			Email: "test2@gmail.com",
		},
		{
			Email: "test3@gmail.com",
		},
		{
			Email: "test4@gmail.com",
		},
		{
			Email: "test5@gmail.com",
		},
	}

	for i, v := range request {
		t.Run(fmt.Sprintf("feat: send otp (handler), test: Success Input %d", i+1), func(t *testing.T) {
			callFunction := userUsecaseMock.Mock.On("SendOTPUsecase", v).Return(nil)

			engine := gin.Default()
			engine.POST("/api/v1/users/otp", userHandler.SendOTP)

			requestDataInJson, err := json.Marshal(v)
			if err != nil {
				t.Fatal(err)
			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", "/api/v1/users/otp", bytes.NewBuffer(requestDataInJson))
			if err != nil {
				t.Fatal(err)
			}
			engine.ServeHTTP(response, request)

			var responseBody map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &responseBody)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, http.StatusOK, response.Code, "http status code should be equal")
			assert.Equal(t, "success", responseBody["status"], "status should be equal")
			assert.Equal(t, "successfully send OTP", responseBody["message"], "message should be equal")
			assert.Nil(t, responseBody["data"], "data should be nil")

			callFunction.Unset()
		})
	}
}

func TestBasicLoginSuccessInput(t *testing.T) {
	request := []domain.BasicLoginBind{
		{
			Email:    "test1@test.com",
			Password: "testpass1",
		},
		{
			Email:    "test2@test.com",
			Password: "testpass2",
		},
		{
			Email:    "test3@test.com",
			Password: "testpass3",
		},
		{
			Email:    "test4@test.com",
			Password: "testpass4",
		},
		{
			Email:    "test5s@test.com",
			Password: "testpass5s",
		},
	}

	for i, v := range request {
		t.Run(fmt.Sprintf("feat: basic login (handler), test: Success Input %d", i+1), func(t *testing.T) {
			functionResponse := struct {
				Token string `json:"token"`
			}{
				"testToken",
			}
			callFunction := userUsecaseMock.Mock.On("BasicLoginUsecase", v).Return(functionResponse, nil)

			engine := gin.Default()
			engine.POST("/api/auth/basic/login", userHandler.BasicLogin)

			requestDataInJson, err := json.Marshal(v)
			if err != nil {
				t.Fatal(err)
			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", "/api/auth/basic/login", bytes.NewBuffer(requestDataInJson))
			if err != nil {
				t.Fatal(err)
			}
			engine.ServeHTTP(response, request)

			var responseBody map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &responseBody)
			if err != nil {
				t.Fatal(err)
			}

			responseData := responseBody["data"].(map[string]any)
			assert.Equal(t, http.StatusOK, response.Code, "http status code should be equal")
			assert.Equal(t, "success", responseBody["status"], "status should be equal")
			assert.Equal(t, "successfully login", responseBody["message"], "message should be equal")
			assert.Equal(t, responseData["token"], "testToken", "token should be equal")

			callFunction.Unset()
		})
	}
}

func TestVerifyForgetPasswordSuccessInput(t *testing.T) {
	request := []domain.VerifyAccountBind{
		{
			Email: "test1@test.com",
			Code:  "1111",
		},
		{
			Email: "test2@test.com",
			Code:  "2111",
		},
		{
			Email: "test3@test.com",
			Code:  "3111",
		},
		{
			Email: "test4@test.com",
			Code:  "4111",
		},
		{
			Email: "test5@test.com",
			Code:  "5111",
		},
	}

	for i, v := range request {
		t.Run(fmt.Sprintf("feat: verify forget password (handler), test: Success Input %d", i+1), func(t *testing.T) {
			apiResponse := struct {
				Token string `json:"token"`
			}{
				"testToken",
			}
			callFunction := userUsecaseMock.Mock.On("VerifyForgetPasswordUsecase", v).Return(apiResponse, nil)

			engine := gin.Default()
			engine.POST("/api/v1/users/verify-fp", userHandler.VerifyForgetPassword)

			requestDataInJson, err := json.Marshal(v)
			if err != nil {
				t.Fatal(err)
			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", "/api/v1/users/verify-fp", bytes.NewBuffer(requestDataInJson))
			if err != nil {
				t.Fatal(err)
			}
			engine.ServeHTTP(response, request)

			var responseBody map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &responseBody)
			if err != nil {
				t.Fatal(err)
			}

			responseData := responseBody["data"].(map[string]any)
			assert.Equal(t, http.StatusOK, response.Code, "http status code should be equal")
			assert.Equal(t, "success", responseBody["status"], "status should be equal")
			assert.Equal(t, "successfully verified forget password", responseBody["message"], "message should be equal")
			assert.Equal(t, responseData["token"], "testToken", "token should be equal")

			callFunction.Unset()
		})
	}
}

func TestForgetPasswordSuccessInput(t *testing.T) {
	request := []domain.ForgetPasswordBind{
		{
			Password:              "testpass1",
			Verification_Password: "testpass1",
		},
		{
			Password:              "testpass2",
			Verification_Password: "testpass2",
		},
		{
			Password:              "testpass3",
			Verification_Password: "testpass3",
		},
		{
			Password:              "testpass4",
			Verification_Password: "testpass4",
		},
		{
			Password:              "testpass5",
			Verification_Password: "testpass5",
		},
	}

	for i, v := range request {
		t.Run(fmt.Sprintf("feat: forget password (handler), test: Success Input %d", i+1), func(t *testing.T) {
			callFunction := userUsecaseMock.Mock.On("ForgetPasswordUsecase", getEmail(), v).Return(nil)

			engine := gin.Default()
			engine.POST("/api/v1/users/forget-password", middlewareSetEmail, userHandler.ForgetPassword)

			requestDataInJson, err := json.Marshal(v)
			if err != nil {
				t.Fatal(err)
			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", "/api/v1/users/forget-password", bytes.NewBuffer(requestDataInJson))
			if err != nil {
				t.Fatal(err)
			}
			engine.ServeHTTP(response, request)

			var responseBody map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &responseBody)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, http.StatusOK, response.Code, "http status code should be equal")
			assert.Equal(t, "success", responseBody["status"], "status should be equal")
			assert.Equal(t, "successfully reset password", responseBody["message"], "message should be equal")
			
			callFunction.Unset()
		})
	}
}

func middlewareSetEmail(c *gin.Context) {
	c.Set("email", "test@test.com")
	c.Next()
}

func getEmail() string {
	return "test@test.com"
}
