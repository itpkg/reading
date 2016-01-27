package site

import (
	"html/template"
	"os"
	"path"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
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

func getHtmlTemplate(names ...string) (*template.Template, error) {
	var files []string
	for _, f := range names {
		files = append(files, path.Join("templates", f+".tmpl"))
	}
	return template.ParseFiles(files...)
}

func writeHtml(tpl *template.Template, id, dist, htm string, model interface{}, mode os.FileMode, min bool) error {

	name := path.Join(dist, "assets", htm)
	dir := path.Dir(name)
	os.MkdirAll(dir, 0755)
	f, err := os.OpenFile(name+".html", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	defer f.Close()

	if min {
		m := minify.New()
		m.AddFunc("text/html", html.Minify)
		m.AddFunc("text/javascript", js.Minify)
		mw := m.Writer("text/html", f)

		return tpl.ExecuteTemplate(mw, id, model)
	} else {
		return tpl.ExecuteTemplate(f, id, model)
	}
}
