package web

import (
	"time"
)

type Response struct {
	Ok        bool        `json:"ok"`
	Payload   interface{} `json:"payload"`
	Messages  []string    `json:"messages"`
	CreatedAt time.Time   `json:"created_at"`
}

func NewResponse(ok bool, payload interface{}, messages ...string) *Response {
	return &Response{
		Ok:        ok,
		Payload:   payload,
		Messages:  messages,
		CreatedAt: time.Now(),
	}
}
