package core

import (
	"github.com/codegangsta/cli"
	"github.com/gorilla/pat"
	"github.com/itpkg/reading/api/rss"
	"github.com/itpkg/reading/api/sitemap"
	"github.com/jinzhu/gorm"
)

type Engine interface {
	Mount(*pat.Router)
	Seed(*gorm.DB) error
	Migrate(*gorm.DB) error
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
