package dto

type PostMessage struct {
	SenderID    string `json:"senderID"`
	RecipientID string `json:"recipientID"`
	Msg         string `json:"msg"`
}
