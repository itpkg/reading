package auth

import (
	"net/http"
)

type Oauth interface {
	Token(req *http.Request) ([]byte, error)
	Url() string
}
