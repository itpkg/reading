package cache

import (
	"net/http"
)

type Provider interface {
	GetOrSet(key string, val interface{}, cb func(interface{}) (uint, error)) error
	Set(key string, val interface{}, minutes uint) error
	Get(key string, val interface{}) error
	Page(wrt http.ResponseWriter, req *http.Request, contentType string, minutes uint, callback func() ([]byte, error))
}
