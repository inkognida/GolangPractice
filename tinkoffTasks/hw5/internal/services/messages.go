package services

import (
	"github.com/sirupsen/logrus"
	"time"
	"tinkoffTasks/hw5/internal/domain"
	"tinkoffTasks/hw5/internal/dto"
)

const (
	shared = "shared"
)

type MessagesRepository interface {
	PostMessage(message domain.Message) error
	GetMessages(message dto.GetMessage) ([]domain.Message, error)
}

type MessagesService struct {
	logger       *logrus.Logger
	messagesRepo MessagesRepository
}

func NewMessagesService(repository MessagesRepository, logger *logrus.Logger) *MessagesService {
	return &MessagesService{
		logger:       logger,
		messagesRepo: repository,
	}
}

func (service *MessagesService) GetMessages(message dto.GetMessage) ([]domain.Message, error) {
	out, err := service.messagesRepo.GetMessages(message)
	if err != nil {
		return []domain.Message{}, err
	}

	return out, nil
}

func (service *MessagesService) PostMessage(message domain.Message) error {
	addMessage := domain.Message{
		SenderID:    message.SenderID,
		RecipientID: message.RecipientID,
		Msg:         message.Msg,
		DateTime:    time.Time{},
	}

	return service.messagesRepo.PostMessage(addMessage)
}
