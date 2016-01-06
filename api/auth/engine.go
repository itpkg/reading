package auth

import (
	"github.com/codegangsta/cli"
	"github.com/gorilla/pat"
	"github.com/itpkg/reading/api/core"
	"github.com/itpkg/reading/api/rss"
	"github.com/itpkg/reading/api/sitemap"
	"golang.org/x/tools/blog/atom"
)

type AuthEngine struct {
}

func (p *AuthEngine) Mount(*pat.Router) {
	//todo
}
func (p *AuthEngine) Seed() {
	//todo
}
func (p *AuthEngine) Migrate() {
	//todo

}

func (p *AuthEngine) Sitemap() sitemap.Handler {
	return func() []*sitemap.Url {
		return []*sitemap.Url{} //todo
	}
}
func (p *AuthEngine) Rss() rss.Handler {
	return func(lang string) []*atom.Entry {
		return []*atom.Entry{} //todo
	}
}

func (p *AuthEngine) Shell() []cli.Command {
	return []cli.Command{}
}

func init() {
	core.Register(&AuthEngine{})
}
