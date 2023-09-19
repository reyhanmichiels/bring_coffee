package usecase

import user_repository "github.com/reyhanmichiels/bring_coffee/app/user/repository"

type IUserUsecase interface {

}

type UserUsecase struct {
	UserRepo user_repository.IUserRepository
}

func NewUserUsecase(userRepo user_repository.IUserRepository) IUserUsecase {
	return &UserUsecase{
		UserRepo: userRepo,
	}
}