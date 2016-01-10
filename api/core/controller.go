package core

import (
	"github.com/unrolled/render"
	"net/http"
)

type Controller struct {
}

func (p *Controller) Pager(r *render.Render, w http.ResponseWriter, pg *Pager, i interface{}) {
	r.JSON(w, http.StatusOK, map[string]interface{}{"pager": pg, "items": i})
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
