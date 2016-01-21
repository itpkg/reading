package token

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type parseFunc func(jwt.Keyfunc) (*jwt.Token, error)

type Provider interface {
	ParseFromRequest(*http.Request) (map[string]interface{}, error)
	Parse(str string) (map[string]interface{}, error)
	New(data map[string]interface{}, minutes uint) (string, error)
}
