package blog

import (
	"net/http"

	"github.com/itpkg/reading/api/core"
	"github.com/julienschmidt/httprouter"
)

func (p *BlogEngine) Mount(rt core.Router) {
	rt.GET("/blog/:id", func(http.ResponseWriter, *http.Request, httprouter.Params) {
		//todo
	})
	rt.GET("/blog", func(http.ResponseWriter, *http.Request, httprouter.Params) {
		//todo
	})
}
