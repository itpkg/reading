package blog

import (
	"github.com/codegangsta/cli"
	"github.com/itpkg/reading/api/core"
	"github.com/itpkg/reading/api/rss"
	"github.com/itpkg/reading/api/sitemap"
	"golang.org/x/tools/blog/atom"
)

type BlogEngine struct {
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
