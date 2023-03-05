package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"net/http"
	"tinkoffTasks/hw5/internal/dto"
	"tinkoffTasks/hw5/pkg/response"
)

const (
	wrongBody = "wrong body request"
)

type UserService interface {
	Register(registration dto.Register) (dto.User, error)
	Login(login dto.Register) (dto.User, error)
	Logout(token string) error
	CheckAuth(token string) (string, error)
}

type UsersHandler struct {
	usersService UserService
	logger       *logrus.Logger
}

func NewUsersHandler(service UserService, logger *logrus.Logger) *UsersHandler {
	return &UsersHandler{
		usersService: service,
		logger:       logger,
	}
}

func (handler *UsersHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/register", handler.register)
	r.Post("/login", handler.login)
	r.Delete("/logout", handler.logout)

	return r
}

func (handler *UsersHandler) logout(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Token")

	if token == "" {
		response.Respond(w, http.StatusUnauthorized, nil)
		return
	}

	err := handler.usersService.Logout(token)
	if err != nil {
		response.Respond(w, http.StatusInternalServerError, dto.Error{Message: err.Error()})
		return
	}

	response.Respond(w, http.StatusOK, nil)
}

func (handler *UsersHandler) register(w http.ResponseWriter, r *http.Request) {
	register := dto.Register{}

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil || register.Login == "" || register.Password == "" {
		response.Respond(w, http.StatusBadRequest, dto.Error{Message: wrongBody})
		return
	}
	defer r.Body.Close()

	user, err := handler.usersService.Register(register)
	if err != nil {
		response.Respond(w, http.StatusInternalServerError,
			dto.Error{Message: err.Error()})
		return
	}

	response.Respond(w, http.StatusOK, user)
}

func (handler *UsersHandler) login(w http.ResponseWriter, r *http.Request) {
	register := dto.Register{}

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil || register.Login == "" || register.Password == "" {
		response.Respond(w, http.StatusBadRequest, dto.Error{Message: wrongBody})
	}
	defer r.Body.Close()

	user, err := handler.usersService.Login(register)
	if err != nil {
		response.Respond(w, http.StatusInternalServerError, dto.Error{Message: err.Error()})
		return
	}

	response.Respond(w, http.StatusOK, user)
}
