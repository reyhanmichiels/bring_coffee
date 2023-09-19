package handler

import user_usecase "github.com/reyhanmichiels/bring_coffee/app/user/usecase"

type UserHandler struct {
	UserUsecase user_usecase.IUserUsecase
}

func NewUserHandler(userUsecase user_usecase.IUserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: userUsecase,
	}
}
