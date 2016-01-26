package storage

type Provider interface {
	Delete(url string) error
	Store(name string, body []byte) (string, error)
}
