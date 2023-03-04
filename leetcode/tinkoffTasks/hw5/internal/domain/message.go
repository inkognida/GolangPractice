package domain

import "time"

type Message struct {
	SenderID    string
	RecipientID string
	Msg         string
	DateTime    time.Time
}
