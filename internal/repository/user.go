package repository

import (
	"UAKI-WEB/entity"
	"UAKI-WEB/model"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	RegisterUser(user entity.User) (entity.User, error)
	GetUser(param model.UserParam) (entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

// GetUser implements UserRepositoryInterface.
func (u *UserRepository) GetUser(param model.UserParam) (entity.User, error) {
	user := entity.User{}
	err := u.db.Debug().Where(&param).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// Create implements UserRepositoryInterface.
func (u *UserRepository) RegisterUser(user entity.User) (entity.User, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{db: db}
}
