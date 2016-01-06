package books

import (
	"net/http"

	"github.com/codegangsta/cli"
	"github.com/itpkg/reading/api/core"
	"github.com/itpkg/reading/api/rss"
	"github.com/itpkg/reading/api/sitemap"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/tools/blog/atom"
)

type BooksEngine struct {
}

func (p *BooksEngine) Mount(rt core.Router) {
	rt.GET("/books/:id", func(http.ResponseWriter, *http.Request, httprouter.Params) {
		//todo
	})
	rt.GET("/books", func(http.ResponseWriter, *http.Request, httprouter.Params) {
		//todo
	})
}
func (p *BooksEngine) Seed() error {
	//todo
	return nil
}
func (p *BooksEngine) Migrate() {
	//todo

}

func (p *BooksEngine) Sitemap() sitemap.Handler {
	return func() []*sitemap.Url {
		return []*sitemap.Url{} //todo
	}
}
func (p *BooksEngine) Rss() rss.Handler {
	return func(lang string) []*atom.Entry {
		return []*atom.Entry{} //todo
	}
}

func (p *BooksEngine) Shell() []cli.Command {
	return []cli.Command{}
}

func init() {
	core.Register(&BooksEngine{})
}
