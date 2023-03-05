package handlers

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"tinkoffTasks/hw5/internal/domain"
	"tinkoffTasks/hw5/internal/dto"
	"tinkoffTasks/hw5/pkg/response"
)

const (
	userIdCtx    = "userID"
	badAuthToken = "Can't use auth token to get UserID"
	shared       = "shared"
)

type MessageService interface {
	GetMessages(message dto.GetMessage) ([]domain.Message, error)
	PostMessage(message domain.Message) error
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

	r.Get("/{userId}", handler.getMessages)
	r.Post("/{userId}", handler.postMessages)

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
	offsetTmp := r.URL.Query().Get("offset")
	if offsetTmp == "" {
		offsetTmp = "0"
	}
	offset, err := strconv.Atoi(offsetTmp)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, dto.Error{Message: err.Error()})
		return
	}

	limitTmp := r.URL.Query().Get("limit")
	if limitTmp == "" {
		limitTmp = "10"
	}
	limit, err := strconv.Atoi(limitTmp)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, dto.Error{Message: err.Error()})
		return
	}
	senderID, ok := r.Context().Value(userIdCtx).(string)
	if !ok {
		response.Respond(w, http.StatusBadRequest, dto.Error{Message: badAuthToken})
		return
	}

	// TODO check URLParam (recipientID)
	recipientID := chi.URLParam(r, "userId")
	if recipientID == "" {
		recipientID = shared
	}

	getMessage := dto.GetMessage{
		SenderID:    senderID,
		RecipientID: recipientID,
		Limit:       limit,
		Offset:      offset,
	}

	messages, err := handler.messageService.GetMessages(getMessage)
	if err != nil {
		response.Respond(w, http.StatusNotFound, dto.Error{Message: err.Error()})
		return
	}

	response.Respond(w, http.StatusOK, messages)
}

func (handler *MessagesHandler) postMessages(w http.ResponseWriter, r *http.Request) {
	senderID, ok := r.Context().Value(userIdCtx).(string)
	if !ok {
		response.Respond(w, http.StatusInternalServerError, dto.Error{Message: badAuthToken})
		return
	}
	recipientID := chi.URLParam(r, "userId")
	if recipientID == "" {
		recipientID = shared
	}
	msg := dto.PostMessage{}
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		response.Respond(w, http.StatusBadRequest, dto.Error{Message: wrongBody})
		return
	}

	postMessage := domain.Message{
		SenderID:    senderID,
		RecipientID: recipientID,
		Msg:         msg.Msg,
	}

	err := handler.messageService.PostMessage(postMessage)
	if err != nil {
		response.Respond(w, http.StatusInternalServerError, dto.Error{Message: err.Error()})
		return
	}

	response.Respond(w, http.StatusOK, nil)
}
