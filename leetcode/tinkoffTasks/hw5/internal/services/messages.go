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
	out := make([]domain.Message, 0)
	var err error

	if message.SenderID == shared {
		out, err = service.messagesRepo.GetSharedMessages(message)
		if err != nil {
			return []domain.Message{}, err
		}

		return out, nil
	} else {
		out, err = service.messagesRepo.GetPrivateMessages(message)
		if err != nil {
			return []domain.Message{}, err
		}

		return out, nil
	}
}

func (service *MessagesService) PostMessage(message dto.PostMessage) error {
	addMessage := domain.Message{
		SenderID:    message.SenderID,
		RecipientID: message.RecipientID,
		Msg:         message.Msg,
		DateTime:    time.Time{},
	}

	return service.messagesRepo.PostMessage(addMessage)
}
