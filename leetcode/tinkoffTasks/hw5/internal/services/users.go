package services

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"tinkoffTasks/hw5/internal/domain"
	"tinkoffTasks/hw5/internal/dto"
	"tinkoffTasks/hw5/pkg/auth"
)

const (
	wrongPassword = "Wrong password for this user"
)

type UsersRepository interface {
	AddUser(login string, passcode string) (domain.User, error)
	UserWithToken(login string, token string) (domain.User, error)
	GetUser(login string) (domain.User, error)
	UserLogout(token string) error
	CheckAuth(token string) (string, error)
}

type UserService struct {
	userRepo UsersRepository
	logger   *logrus.Logger
}

func NewUsersService(repository UsersRepository, logger *logrus.Logger) *UserService {
	return &UserService{
		userRepo: repository,
		logger:   logger,
	}
}

func (service *UserService) Register(registration dto.Register) (dto.User, error) {
	userLogin := registration.Login
	userPassword, err := auth.HashPassword(registration.Password)
	if err != nil {
		return dto.User{}, err
	}

	user, err := service.userRepo.AddUser(userLogin, userPassword)
	if err != nil {
		return dto.User{}, err
	}

	return dto.User{
		ID:    user.ID,
		Login: user.Login,
	}, nil
}

func (service *UserService) Login(login dto.Register) (dto.User, error) {
	user, err := service.userRepo.GetUser(login.Login)
	if err != nil {
		return dto.User{}, err
	}

	if !auth.CheckPasswordHash(login.Password, user.Password) {
		return dto.User{}, fmt.Errorf("%s", wrongPassword)
	}

	token, err := auth.GenerateToken(login.Login)
	if err != nil {
		return dto.User{}, err
	}

	user, err = service.userRepo.UserWithToken(login.Login, token)
	if err != nil {
		return dto.User{}, err
	}

	return dto.User{
		ID:    user.ID,
		Login: user.Login,
		Token: user.Token,
	}, err
}

func (service *UserService) Logout(token string) error {
	return service.userRepo.UserLogout(token)
}

func (service *UserService) CheckAuth(token string) (string, error) {
	return service.userRepo.CheckAuth(token)
}
