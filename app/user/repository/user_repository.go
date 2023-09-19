package repository

import (
	"github.com/reyhanmichiels/bring_coffee/domain"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user *domain.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (userRepo *UserRepository) CreateUser(user *domain.User) error {
	err := userRepo.db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}