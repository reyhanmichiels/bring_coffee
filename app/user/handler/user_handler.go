package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	user_usecase "github.com/reyhanmichiels/bring_coffee/app/user/usecase"
	"github.com/reyhanmichiels/bring_coffee/domain"
	"github.com/reyhanmichiels/bring_coffee/util"
)

type UserHandler struct {
	UserUsecase user_usecase.IUserUsecase
}

func NewUserHandler(userUsecase user_usecase.IUserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: userUsecase,
	}
}

func (userHandler *UserHandler) Registration(c *gin.Context) {
	var request domain.RegistBind
	err := c.ShouldBindJSON(&request)
	if err != nil {
		util.FailedResponse(c, http.StatusBadRequest, "failed bind input data", err)
		return
	}

	apiData, errObject := userHandler.UserUsecase.RegistrationUsecase(request)
	if errObject != nil {
		errObject := errObject.(util.ErrorObject)
		util.FailedResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return
	}

	util.SuccessedResponse(c, http.StatusCreated, "successfully registration new user", apiData)
}

func (userHandler *UserHandler) VerifyAccount(c *gin.Context) {
	var request domain.VerifyAccountBind
	err := c.ShouldBindJSON(&request)
	if err != nil {
		util.FailedResponse(c, http.StatusBadRequest, "failed bind input data", err)
		return
	}

	errObject := userHandler.UserUsecase.VerifyAccountUsecase(request)
	if errObject != nil {
		errObject := errObject.(util.ErrorObject)
		util.FailedResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return
	}

	util.SuccessedResponse(c, http.StatusOK, "successfully verified account", nil)
}

func (userHandler *UserHandler) SendOTP(c *gin.Context) {
	var request domain.SendOTPBind
	err := c.ShouldBindJSON(&request)
	if err != nil {
		util.FailedResponse(c, http.StatusBadRequest, "failed bind input data", err)
		return
	}

	errObject := userHandler.UserUsecase.SendOTPUsecase(request)
	if errObject != nil {
		errObject := errObject.(util.ErrorObject)
		util.FailedResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return
	}

	util.SuccessedResponse(c, http.StatusOK, "successfully send OTP", nil)
}

func (userHandler *UserHandler) BasicLogin(c *gin.Context) {
	var request domain.BasicLoginBind
	err := c.ShouldBindJSON(&request)
	if err != nil {
		util.FailedResponse(c, http.StatusBadRequest, "failed bind input data", err)
		return
	}

	apiData, errObject := userHandler.UserUsecase.BasicLoginUsecase(request)
	if errObject != nil {
		errObject := errObject.(util.ErrorObject)
		util.FailedResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return
	}

	util.SuccessedResponse(c, http.StatusOK, "successfully login", apiData)
}

func (userHandler *UserHandler) VerifyForgetPassword(c *gin.Context) {
	var request domain.VerifyAccountBind
	err := c.ShouldBindJSON(&request)
	if err != nil {
		util.FailedResponse(c, http.StatusBadRequest, "failed bind input data", err)
		return
	}

	apiData, errObject := userHandler.UserUsecase.VerifyForgetPasswordUsecase(request)
	if errObject != nil {
		errObject := errObject.(util.ErrorObject)
		util.FailedResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return
	}

	util.SuccessedResponse(c, http.StatusOK, "successfully verified forget password", apiData)
}

func (userHandler *UserHandler) ForgetPassword(c *gin.Context) {
	var request domain.ForgetPasswordBind
	err := c.ShouldBindJSON(&request)
	if err != nil {
		util.FailedResponse(c, http.StatusBadRequest, "failed bind input data", err)
		return
	}

	email, ok := c.Get("email")
	if !ok {
		util.FailedResponse(c, http.StatusInternalServerError, "user email not found", errors.New(""))
	}

	errObject := userHandler.UserUsecase.ForgetPasswordUsecase(email.(string), request)
	if errObject != nil {
		errObject := errObject.(util.ErrorObject)
		util.FailedResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return
	}

	util.SuccessedResponse(c, http.StatusOK, "successfully reset password", nil)
}
