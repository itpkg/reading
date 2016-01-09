package assets

import (
	"github.com/itpkg/reading/api/cache"
	"github.com/itpkg/reading/api/core"
	"github.com/itpkg/reading/api/rss"
	"github.com/itpkg/reading/api/sitemap"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/unrolled/render"
	"golang.org/x/tools/blog/atom"
)

type AssetsEngine struct {
	core.Controller

	Render *render.Render  `inject:""`
	Db     *gorm.DB        `inject:""`
	Logger *logging.Logger `inject:""`
	Cache  cache.Provider  `inject:""`
}

func (p *AssetsEngine) Seed() error {
	return nil
}

func (p *AssetsEngine) Sitemap() sitemap.Handler {
	return func() []*sitemap.Url {
		return []*sitemap.Url{} //todo
	}
}
func (p *AssetsEngine) Rss() rss.Handler {
	return func(lang string) []*atom.Entry {
		return []*atom.Entry{}
	}
}

func init() {
	core.Register(&AssetsEngine{})
}
