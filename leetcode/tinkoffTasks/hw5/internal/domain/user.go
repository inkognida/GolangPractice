package domain

import "github.com/gofrs/uuid"

type User struct {
	ID       uuid.UUID
	Login    string
	Password string
	Nick     string
	Token    string
}
