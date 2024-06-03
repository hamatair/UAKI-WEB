package service

import (
	"UAKI-WEB/internal/repository"
	"UAKI-WEB/pkg/bcrypt"
	"UAKI-WEB/pkg/jwt"
)

type Service struct {
	UserService UserServiceInterface
}

type InitParam struct {
	Repository *repository.Repository
	JwtAuth jwt.Interface
	Bcrypt bcrypt.Interface
}

func NewService(param InitParam) *Service {
	return &Service{
		UserService: NewUserService(param.Repository.UserRepository, param.Bcrypt, param.JwtAuth),
	}
}
