package handlers

import "github.com/go-chi/chi/v5"

type UserService interface {
	Register()
	Login()
}

type User struct {
	service UserService
}

func NewUser(service UserService) *User {
	return &User{service: service}
}

func (u *User) Routes() chi.Router {
	r := chi.NewRouter()
}
