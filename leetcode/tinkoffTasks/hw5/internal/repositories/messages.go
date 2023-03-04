package repositories

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"sync"
	"tinkoffTasks/hw5/internal/domain"
	"tinkoffTasks/hw5/internal/dto"
)

const (
	shared           = "shared"
	noSharedMessages = "There is no messages in shared chat"
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

func (repo *MessagesRepository) GetSharedMessages(message dto.GetMessage) ([]domain.Message, error) {
	out := make([]domain.Message, 0)
	repo.messagesMutex.RLock()
	defer repo.messagesMutex.RUnlock()

	for i := message.Offset; i < message.Offset+message.Limit && i < len(repo.messages); i++ {
		if repo.messages[i].RecipientID == shared {
			out = append(out, repo.messages[i])
		}
	}

	if len(out) == 0 {
		return []domain.Message{}, fmt.Errorf("%s", noSharedMessages)
	}

	return out, nil
}

func (repo *MessagesRepository) GetPrivateMessages(message dto.GetMessage) ([]domain.Message, error) {
	out := make([]domain.Message, 0)
	repo.messagesMutex.RLock()
	defer repo.messagesMutex.RUnlock()

	for i := message.Offset; i < message.Offset+message.Limit && i < len(repo.messages); i++ {
		if repo.messages[i].SenderID == message.SenderID &&
			repo.messages[i].RecipientID == message.RecipientID { //TODO check || condition
			out = append(out, repo.messages[i])
		}
	}

	return out, nil
}

func (repo *MessagesRepository) PostMessage(message domain.Message) error {
	repo.messagesMutex.Lock()
	repo.messages = append(repo.messages, message)
	repo.messagesMutex.Unlock()

	return nil
}

func (repo *MessagesRepository) CheckAuth(token string) error {
	repo.messagesMutex.RLock()
	for _,
}
