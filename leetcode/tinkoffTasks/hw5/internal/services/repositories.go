package services

import (
	"tinkoffTasks/hw5/internal/domain"
	"tinkoffTasks/hw5/internal/dto"
)

type UsersRepository interface {
	AddUser(login string, passcode string) (domain.User, error)
	AddUserToken(login string, token string) domain.User
	GetUser(login string) (domain.User, error)
	UserLogout(token string) error
	CheckAuth(token string) (string, error)
}

type MessagesRepository interface {
	GetSharedMessages(message dto.GetMessage) ([]domain.Message, error)
	GetPrivateMessages(message dto.GetMessage) ([]domain.Message, error)
	PostMessage(message domain.Message) error
}
