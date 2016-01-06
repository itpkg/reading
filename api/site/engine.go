package site

import (
	"github.com/gorilla/pat"
	"github.com/itpkg/reading/api/core"
	"github.com/itpkg/reading/api/rss"
	"github.com/itpkg/reading/api/sitemap"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"golang.org/x/tools/blog/atom"
)

type SiteEngine struct {
	Db     *gorm.DB        `inject:""`
	Dao    *Dao            `inject:""`
	Logger *logging.Logger `inject:""`
}

func (p *SiteEngine) Mount(*pat.Router) {
}

//=========================================================
func (p *SiteEngine) Seed() {
	//todo

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
