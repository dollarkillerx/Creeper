package models

type Message struct {
	Index string `json:"-"`

	ID             string `json:"id"`
	Message        string `json:"message"`
	CreateAt       int64  `json:"create_at"`
	CreateAtString string `json:"create_at_string"`
}
