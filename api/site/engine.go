package site

import (
	"github.com/itpkg/reading/api/core"
	"github.com/itpkg/reading/api/rss"
	"github.com/itpkg/reading/api/sitemap"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/unrolled/render"
	"golang.org/x/tools/blog/atom"
)

type SiteEngine struct {
	core.Controller

	Db     *gorm.DB        `inject:""`
	Dao    *Dao            `inject:""`
	Logger *logging.Logger `inject:""`
	Render *render.Render  `inject:""`
}

func (p *SiteEngine) Mount(rt core.Router) {
	rt.GET("/info", p.info)
	rt.GET("/baidu_verify_:id", p.baidu)
	rt.GET("/google:id", p.google)
	rt.GET("/rss.atom", p.rss)
	rt.GET("/sitemap.xml", p.sitemap)
	rt.GET("/robots.txt", p.robots)
}

//=========================================================
func (p *SiteEngine) Seed() error {
	//todo
	return nil

}

func (p *SiteEngine) Migrate() {
	db := p.Db
	db.AutoMigrate(&Locale{}, &Setting{})
	db.Model(&Locale{}).AddUniqueIndex("idx_locales_code_lang", "code", "lang")
}

func (p *SiteEngine) Sitemap() sitemap.Handler {
	return func() []*sitemap.Url {
		return []*sitemap.Url{} //todo
	}
}

func (p *SiteEngine) Rss() rss.Handler {
	return func(lang string) []*atom.Entry {
		return []*atom.Entry{} //todo
	}
}

func init() {
	core.Register(&SiteEngine{})
}
