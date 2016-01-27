package cms

import (
	"fmt"

	"github.com/itpkg/reading/api/auth"
	"github.com/itpkg/reading/api/cache"
	"github.com/itpkg/reading/api/core"
	"github.com/itpkg/reading/api/rss"
	"github.com/itpkg/reading/api/site"
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
	SiteDao *site.Dao        `inject:""`
}

func (p *CmsEngine) Asserts() []*core.Template {
	var tps []*core.Template
	var tags []Tag
	p.Db.Select("name").Order("updated_at DESC").Find(&tags)

	for _, lang := range p.SiteDao.Languages() {
		tps = append(tps, &core.Template{
			Lang:    lang,
			Tpl:     "cms_tags",
			Htm:     fmt.Sprintf("cms/tags-%s", lang),
			Payload: tags,
		})
	}
	var articles []Article
	p.Db.Find(&articles)
	for _, a := range articles {
		tps = append(tps, &core.Template{
			Lang:    a.Lang,
			Tpl:     "cms_article",
			Htm:     fmt.Sprintf("cms/article-%02d/%s", a.ID%100, a.Aid),
			Payload: a,
		})
	}

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
