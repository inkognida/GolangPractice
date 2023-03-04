package dto

type User struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	Token string `json:"token,omitempty"`
}
