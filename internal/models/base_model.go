package models

type Message struct {
	Message  string `json:"message"`
	CreateAt int64  `json:"create_at"`
}
