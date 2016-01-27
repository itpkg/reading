package site

import (
	"html/template"
	"os"
	"path"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
)

type SiteModel struct {
	Lang        string
	Title       string
	SubTitle    string
	Keywords    string
	Description string
	AuthorName  string
	AuthorEmail string
	Copyright   string
}

func writeHtml(tpl, dist, htm string, model interface{}, mode os.FileMode, min bool) error {

	t, err := template.ParseFiles(
		path.Join("templates", "layout.tmpl"),
		path.Join("templates", tpl+".tmpl"))
	if err != nil {
		return err
	}
	name := path.Join(dist, "assets", htm)
	dir := path.Dir(name)
	os.MkdirAll(dir, 0700)
	f, err := os.OpenFile(name+".html", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	defer f.Close()

	if min {
		m := minify.New()
		m.AddFunc("text/html", html.Minify)
		mw := m.Writer("text/html", f)

		return t.Execute(mw, model)
	} else {
		return t.Execute(f, model)
	}
}
