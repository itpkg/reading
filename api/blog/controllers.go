package blog

import (
	"github.com/itpkg/reading/api/core"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (p *BlogEngine) index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func (p *BlogEngine) show(http.ResponseWriter, *http.Request, httprouter.Params) {

}

func (p *BlogEngine) Mount(rt core.Router) {
	rt.GET("/blog/:id", p.show)
	rt.GET("/blog", p.index)
}
