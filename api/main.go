package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"net/http"
)

func Test1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "vars, %v\n", mux.Vars(r))
	fmt.Fprintf(w, "query, %v\n", r.URL.Query())
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	r := pat.New()
	r.Get("/test1/{k}/{f}.{i}", Test1)

	http.ListenAndServe(":8080", r)
}
