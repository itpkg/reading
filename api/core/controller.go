package core

import (
	"net/http"
)

type Controller struct {
}

func (p *Controller) NotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

func (p *Controller) Locale(r *http.Request) string {
	l := r.URL.Query().Get("locale")
	if l == "" {
		l = "en-US"
	}
	return l
}
