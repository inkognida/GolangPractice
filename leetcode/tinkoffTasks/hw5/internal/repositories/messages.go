package repositories

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"sync"
	"tinkoffTasks/hw5/internal/domain"
	"tinkoffTasks/hw5/internal/dto"
)

const (
	noMessages = "There is no messages for you"
)

type MessagesRepository struct {
	logger        *logrus.Logger
	messages      []domain.Message
	messagesMutex *sync.RWMutex
}

func NewMessagesRepository(logger *logrus.Logger) *MessagesRepository {
	return &MessagesRepository{
		logger:        logger,
		messages:      make([]domain.Message, 0),
		messagesMutex: &sync.RWMutex{},
	}
}

func (repo *MessagesRepository) GetMessages(message dto.GetMessage) ([]domain.Message, error) {
	out := make([]domain.Message, 0)
	repo.messagesMutex.RLock()
	defer repo.messagesMutex.RUnlock()

	for i := message.Offset; i < message.Offset+message.Limit && i < len(repo.messages); i++ {
		if repo.messages[i].RecipientID == message.RecipientID {
			out = append(out, repo.messages[i])
		}
	}

	if len(out) == 0 {
		return []domain.Message{}, fmt.Errorf("%s", noMessages)
	}

	return out, nil
}

func (repo *MessagesRepository) PostMessage(message domain.Message) error {
	repo.messagesMutex.Lock()
	repo.messages = append(repo.messages, message)
	repo.messagesMutex.Unlock()

	return nil
}
