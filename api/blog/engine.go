package blog

import (
	"net/http"

	"github.com/codegangsta/cli"
	"github.com/itpkg/reading/api/core"
	"github.com/itpkg/reading/api/rss"
	"github.com/itpkg/reading/api/sitemap"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/tools/blog/atom"
)

type BlogEngine struct {
}

func (p *BlogEngine) Mount(rt core.Router) {
	rt.GET("/blog/:id", func(http.ResponseWriter, *http.Request, httprouter.Params) {
		//todo
	})
	rt.GET("/blog", func(http.ResponseWriter, *http.Request, httprouter.Params) {
		//todo
	})
}
func (p *BlogEngine) Seed() error {
	//todo
	return nil
}
func (p *BlogEngine) Migrate() {
	//todo

}

func (p *BlogEngine) Sitemap() sitemap.Handler {
	return func() []*sitemap.Url {
		return []*sitemap.Url{} //todo
	}
}
func (p *BlogEngine) Rss() rss.Handler {
	return func(lang string) []*atom.Entry {
		return []*atom.Entry{} //todo
	}
}

func (p *BlogEngine) Shell() []cli.Command {
	return []cli.Command{}
}

func init() {
	core.Register(&BlogEngine{})
}
