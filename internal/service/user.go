package service

import (
	"UAKI-WEB/entity"
	"UAKI-WEB/internal/repository"
	"UAKI-WEB/model"
	"UAKI-WEB/pkg/bcrypt"
	"UAKI-WEB/pkg/jwt"

	"github.com/google/uuid"
)

type UserServiceInterface interface {
	RegisterUser(param model.RegisterUser) (entity.User, error)
	GetUser(param model.UserParam) (entity.User, error)
}

type UserService struct {
	userRepository repository.UserRepositoryInterface
	bcrypt         bcrypt.Interface
	jwtAuth        jwt.Interface
}

func NewUserService(repository repository.UserRepositoryInterface, bcrypt bcrypt.Interface, jwtAuth jwt.Interface) UserServiceInterface {
	return &UserService{
		userRepository: repository,
		bcrypt: bcrypt,
		jwtAuth: jwtAuth,
	}
}

// GetUser implements UserServiceInterface.
func (u *UserService) GetUser(param model.UserParam) (entity.User, error) {
	return u.userRepository.GetUser(param)

}

// Create implements UserServiceInterface.
func (u *UserService) RegisterUser(param model.RegisterUser) (entity.User, error) {
	hashPassword, err := u.bcrypt.GenerateFromPassword(param.Password)
	if err != nil {
		return entity.User{}, err
	}

	NewUser := entity.User{
		ID: uuid.New(),
		UserName: param.UserName,
		Email: param.Email,
		Password: hashPassword,
	}

	newUser, err := u.userRepository.RegisterUser(NewUser)
	if err != nil {
		return entity.User{}, err
	}

	return newUser, err
}