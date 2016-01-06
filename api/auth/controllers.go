package auth

import (
	"net/http"

	"github.com/itpkg/reading/api/core"
	"github.com/julienschmidt/httprouter"
)

func (p *AuthEngine) Mount(rt core.Router) {
	rt.POST("/users/sign_in", func(http.ResponseWriter, *http.Request, httprouter.Params) {
		//todo
	})
	rt.POST("/users/sign_up", func(http.ResponseWriter, *http.Request, httprouter.Params) {
		//todo
	})
	rt.POST("/users/confirm", func(http.ResponseWriter, *http.Request, httprouter.Params) {
		//todo
	})
	rt.DELETE("/users/sign_out", func(http.ResponseWriter, *http.Request, httprouter.Params) {
		//todo
	})
	rt.POST("/users/forgot_password", func(http.ResponseWriter, *http.Request, httprouter.Params) {
		//todo
	})
	rt.POST("/users/reset_password", func(http.ResponseWriter, *http.Request, httprouter.Params) {
		//todo
	})
	rt.POST("/users/unlock", func(http.ResponseWriter, *http.Request, httprouter.Params) {
		//todo
	})
}
