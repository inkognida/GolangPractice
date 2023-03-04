package handlers

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"tinkoffTasks/hw5/internal/domain"
	"tinkoffTasks/hw5/internal/dto"
	"tinkoffTasks/hw5/pkg/response"
)

const (
	userIdCtx = "userID"
)

type MessageService interface {
	GetMessages(message dto.GetMessage) ([]domain.Message, error)
	PostMessage(message dto.PostMessage) error
}

type MessagesHandler struct {
	messageService MessageService
	usersService   UserService
	logger         *logrus.Logger
}

func NewMessagesHandler(service MessageService, usersService UserService,
	logger *logrus.Logger) *MessagesHandler {
	return &MessagesHandler{
		messageService: service,
		usersService:   usersService,
		logger:         logger,
	}
}

func (handler *MessagesHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Use(handler.authMiddleware)
	r.Get("/", handler.getMessages)
	r.Post("/", handler.postMessages)

	return r
}

func (handler *MessagesHandler) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		token := request.Header.Get("Token")

		userId, err := handler.usersService.CheckAuth(token)
		if err != nil {
			response.Respond(writer, http.StatusUnauthorized, nil)
			return
		}

		ctx := context.WithValue(request.Context(), userIdCtx, userId)
		request = request.WithContext(ctx)
		next.ServeHTTP(writer, request)
	})
}

func (handler *MessagesHandler) getMessages(w http.ResponseWriter, r *http.Request) {
	getMessage := dto.GetMessage{}

	offset := r.URL.Query().Get("offset")
	if offset == "" {
		offset = "0"
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, dto.Error{Message: err.Error()})
		return
	}
	getMessage.Offset = offsetInt

	limit := r.URL.Query().Get("limit")
	if limit == "" {
		limit = "10"
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, dto.Error{Message: err.Error()})
		return
	}
	getMessage.Limit = limitInt

}

func (handler *MessagesHandler) postMessages(w http.ResponseWriter, r *http.Request) {

}
