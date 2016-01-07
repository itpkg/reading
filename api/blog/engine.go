package blog

import (
	"github.com/codegangsta/cli"
	"github.com/itpkg/reading/api/cache"
	"github.com/itpkg/reading/api/core"
	"github.com/itpkg/reading/api/rss"
	"github.com/itpkg/reading/api/sitemap"
	"github.com/op/go-logging"
	"github.com/unrolled/render"
	"golang.org/x/tools/blog/atom"
)

type BlogEngine struct {
	core.Controller

	Render *render.Render  `inject:""`
	Dao    *Dao            `inject:""`
	Logger *logging.Logger `inject:""`
	Cache  cache.Provider  `inject:""`
}

func (p *BlogEngine) Seed() error {
	return nil
}

func (p *BlogEngine) Migrate() {

}

func (p *BlogEngine) Sitemap() sitemap.Handler {
	return func() []*sitemap.Url {
		return []*sitemap.Url{} //todo
	}
}
func (p *BlogEngine) Rss() rss.Handler {
	return func(lang string) []*atom.Entry {
		return []*atom.Entry{}
	}
}

func (p *BlogEngine) Shell() []cli.Command {
	return []cli.Command{}
}

func init() {
	core.Register(&BlogEngine{})
}
