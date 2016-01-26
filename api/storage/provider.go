package storage

import (
	"io"
)

type Provider interface {
	Delete(url string) error
	Store(name string, reader io.Reader) (string, uint, error)
}
