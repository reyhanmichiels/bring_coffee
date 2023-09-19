package handler

import (
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
