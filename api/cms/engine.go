package cms

import (
	"github.com/itpkg/reading/api/auth"
	"github.com/itpkg/reading/api/cache"
	"github.com/itpkg/reading/api/core"
	"github.com/itpkg/reading/api/rss"
	"github.com/itpkg/reading/api/sitemap"
	"github.com/itpkg/reading/api/storage"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/unrolled/render"
	"golang.org/x/tools/blog/atom"
)

type CmsEngine struct {
	core.Controller

	Render  *render.Render   `inject:""`
	Db      *gorm.DB         `inject:""`
	Logger  *logging.Logger  `inject:""`
	Cache   cache.Provider   `inject:""`
	Session *auth.Session    `inject:""`
	AuthDao *auth.Dao        `inject:""`
	Storage storage.Provider `inject:""`
}

func (p *CmsEngine) Asserts() []*core.Template {
	var tps []*core.Template
	return tps
}
func (p *CmsEngine) Seed() error {
	return nil
}

func (p *CmsEngine) Sitemap() sitemap.Handler {
	return func() []*sitemap.Url {
		return []*sitemap.Url{} //todo
	}
}
func (p *CmsEngine) Rss() rss.Handler {
	return func(lang string) []*atom.Entry {
		return []*atom.Entry{}
	}
}

func init() {
	core.Register(&CmsEngine{})
}
