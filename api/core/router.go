package core

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Router interface {
	GET(string, httprouter.Handle)
	POST(string, httprouter.Handle)
	DELETE(string, httprouter.Handle)
	PUT(string, httprouter.Handle)
	PATCH(string, httprouter.Handle)
	ServeFiles(string, http.FileSystem)
}
