package repository

import (
	"github.com/reyhanmichiels/bring_coffee/domain"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	ActivateAccount(userEmail string) error
	FindUserByCondition(user domain.User, condition string, value interface{}, column []string) (domain.User, error)
	Update(user *domain.User, userUpdateData interface{}) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (userRepo *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	tx := userRepo.db.Begin()

	err := tx.Create(user).Error
	if err != nil {
		tx.Rollback()
		return &domain.User{}, err
	}

	return user, tx.Commit().Error
}

func (userRepo *UserRepository) ActivateAccount(userEmail string) error {
	tx := userRepo.db.Begin()

	err := tx.Model(&domain.User{}).Where("email = ?", userEmail).Update("is_verified", true).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (userRepo *UserRepository) FindUserByCondition(user domain.User, condition string, value interface{}, column []string) (domain.User, error) {
	// err := userRepo.db.Model(domain.User{}).First(user, condition, value).Error
	err := userRepo.db.Select(column).First(user, condition, value).Error
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{}, nil
}

func (userRepo *UserRepository) Update(user *domain.User, userUpdateData interface{}) error {
	tx := userRepo.db.Begin()

	err := tx.Model(user).Updates(userUpdateData).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
