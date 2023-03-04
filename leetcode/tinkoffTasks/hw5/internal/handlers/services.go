package handlers

import (
	"tinkoffTasks/hw5/internal/dto"
)

type UserService interface {
	Register(registration dto.Register) (dto.User, error)
	Login(login dto.Register) (dto.User, error)
	Logout(token string) error
	CheckAuth(token string) (string, error)
}
