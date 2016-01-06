package core

import (
	"github.com/codegangsta/cli"
	"github.com/itpkg/reading/api/rss"
	"github.com/itpkg/reading/api/sitemap"
)

type Engine interface {
	Mount(Router)
	Seed() error
	Migrate()
	Sitemap() sitemap.Handler
	Rss() rss.Handler
	Shell() []cli.Command
}

var engines = make([]Engine, 0)

func Register(ens ...Engine) {
	engines = append(engines, ens...)
}

func Loop(fn func(Engine) error) error {
	for _, en := range engines {
		if err := fn(en); err != nil {
			return err
		}
	}
	return nil
}
