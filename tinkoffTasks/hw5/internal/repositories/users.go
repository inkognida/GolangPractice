package repositories

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"sync"
	"tinkoffTasks/hw5/internal/domain"
)

const (
	userExists      = "User with that login is already exist"
	noSuchUser      = "User with that login does not exist"
	noSuchUserToken = "User with incorrect token"
)

type UsersRepository struct {
	logger     *logrus.Logger
	users      []domain.User
	usersMutex *sync.RWMutex
}

func NewUsersRepository(logger *logrus.Logger) *UsersRepository {
	return &UsersRepository{
		logger:     logger,
		users:      make([]domain.User, 0),
		usersMutex: &sync.RWMutex{},
	}
}

func (repo *UsersRepository) AddUser(login string, password string) (domain.User, error) {
	repo.usersMutex.RLock()
	for _, user := range repo.users {
		if user.Login == login {
			repo.usersMutex.RUnlock()
			return domain.User{}, fmt.Errorf("%s", userExists)
		}
	}
	repo.usersMutex.RUnlock()

	user := domain.User{
		ID:       uuid.New().String(),
		Login:    login,
		Password: password,
	}

	repo.usersMutex.Lock()
	repo.users = append(repo.users, user)
	repo.usersMutex.Unlock()

	return user, nil
}

func (repo *UsersRepository) UserWithToken(login string, token string) (domain.User, error) {
	repo.usersMutex.RLock()
	for i := 0; i < len(repo.users); i++ {
		if repo.users[i].Login == login {
			tmp := &repo.users[i]
			tmp.Token = token
			repo.usersMutex.RUnlock()
			return repo.users[i], nil
		}
	}

	repo.usersMutex.RUnlock()
	return domain.User{}, fmt.Errorf("%s", noSuchUser)
}

func (repo *UsersRepository) GetUser(login string) (domain.User, error) {
	repo.usersMutex.RLock()

	for _, user := range repo.users {
		if user.Login == login {
			repo.usersMutex.RUnlock()
			return user, nil
		}
	}

	repo.usersMutex.RUnlock()
	return domain.User{}, fmt.Errorf("%s", noSuchUser)
}

func (repo *UsersRepository) UserLogout(token string) error {
	repo.usersMutex.RLock()

	for i := 0; i < len(repo.users); i++ {
		if repo.users[i].Token == token {
			repo.usersMutex.Lock()
			repo.users = append(repo.users[:i], repo.users[i+1:]...)
			repo.usersMutex.Unlock()
			repo.usersMutex.RUnlock()
			return nil
		}
	}

	repo.usersMutex.RUnlock()
	return fmt.Errorf("%s", noSuchUserToken)
}

func (repo *UsersRepository) CheckAuth(token string) (string, error) {
	repo.usersMutex.RLock()

	for _, user := range repo.users {
		if user.Token == token {
			repo.usersMutex.RUnlock()
			return user.ID, nil
		}
	}

	repo.usersMutex.RUnlock()
	return "", fmt.Errorf("%s", noSuchUserToken)
}