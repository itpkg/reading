package site

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type route struct {
	method string
	path   string
	handle httprouter.Handle
}

type router struct {
	routes []*route
}

func (r *router) GET(p string, h httprouter.Handle) {
	r.add("GET", p, h)
}

func (r *router) POST(p string, h httprouter.Handle) {
	r.add("POST", p, h)
}
func (r *router) DELETE(p string, h httprouter.Handle) {
	r.add("DELETE", p, h)
}
func (r *router) PUT(p string, h httprouter.Handle) {
	r.add("PUT", p, h)
}
func (r *router) PATCH(p string, h httprouter.Handle) {
	r.add("PATCH", p, h)
}
func (r *router) ServeFiles(p string, _ http.FileSystem) {
	r.GET(p, nil)
}
func (r *router) add(m string, p string, h httprouter.Handle) {
	r.routes = append(r.routes, &route{method: m, path: p, handle: h})
}
