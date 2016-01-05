package core

import (
	"github.com/codegangsta/cli"
	"github.com/gorilla/feeds"
	"github.com/gorilla/pat"
	"github.com/jinzhu/gorm"
)

type Engine interface {
	Mount(*pat.Router)
	Seed(*gorm.DB)
	Migrate(*gorm.DB)
	Sitemap()
	Rss() []*feeds.Item
	Shell() []cli.Command
}

var engines = make([]Engine, 0)

func Register(ens ...Engine) {
	engines = append(engines, ens...)
}
