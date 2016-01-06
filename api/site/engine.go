package site

import (
	"net/http"

	"github.com/itpkg/reading/api/core"
	"github.com/itpkg/reading/api/rss"
	"github.com/itpkg/reading/api/sitemap"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"github.com/op/go-logging"
	"golang.org/x/tools/blog/atom"
)

type SiteEngine struct {
	Db     *gorm.DB        `inject:""`
	Dao    *Dao            `inject:""`
	Logger *logging.Logger `inject:""`
}

func (p *SiteEngine) Mount(rt core.Router) {
	rt.GET("/baidu_:id", func(http.ResponseWriter, *http.Request, httprouter.Params) {
		//todo
	})
	rt.GET("/google_:id", func(http.ResponseWriter, *http.Request, httprouter.Params) {
		//todo
	})
	rt.GET("/rss.atom", func(http.ResponseWriter, *http.Request, httprouter.Params) {
		//todo
	})
	rt.GET("/sitemap.xml", func(http.ResponseWriter, *http.Request, httprouter.Params) {
		//todo
	})
	rt.GET("/robots.txt", func(http.ResponseWriter, *http.Request, httprouter.Params) {
		//todo
	})
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
