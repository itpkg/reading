package core

import (
	"net/http"

	"github.com/itpkg/reading/api/cache"
)

type Controller struct {
}

func (p *Controller) Abort(w http.ResponseWriter, e error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(e.Error()))
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
func (p *Controller) CachePage(w http.ResponseWriter, c cache.Provider, k string, t string, m uint, h func() ([]byte, error)) {
	var body []byte
	if err := c.Get(k, &body); err != nil {
		if body, err = h(); err != nil {
			p.Abort(w, err)
			return
		}
		c.Set(k, body, m)

		w.Header().Set("Content-Type", t)
		w.WriteHeader(http.StatusOK)
	} else {
		w.Header().Set("Content-Type", t)
		w.WriteHeader(http.StatusNotModified)
	}
	w.Write(body)
}
