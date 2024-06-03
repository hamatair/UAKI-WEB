package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository UserRepositoryInterface
}

func NewRepository(db *gorm.DB) *Repository{
	return &Repository{
		UserRepository: NewUserRepository(db),
	}
}