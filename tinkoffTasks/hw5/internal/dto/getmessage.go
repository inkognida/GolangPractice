package dto

type GetMessage struct {
	SenderID    string `json:"senderID,omitempty"`
	RecipientID string `json:"recipientID,omitempty"`
	Limit       int    `json:"limit,omitempty"`
	Offset      int    `json:"offset,omitempty"`
}
