package site

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/itpkg/reading/api/cache"
	"github.com/itpkg/reading/api/config"
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

	Cfg    *config.Model   `inject:""`
	Db     *gorm.DB        `inject:""`
	Dao    *Dao            `inject:""`
	Logger *logging.Logger `inject:""`
	Render *render.Render  `inject:""`
	Cache  cache.Provider  `inject:""`
}

//=========================================================
func (p *SiteEngine) Seed() error {
	lf, err := os.Open("tmp/locales.txt")
	if err != nil {
		return err
	}
	san := bufio.NewScanner(lf)
	for san.Scan() {
		line := san.Text()
		ldx := strings.Index(line, ".")
		cdx := strings.Index(line, "=")
		if ldx == -1 || cdx == -1 {
			return errors.New(fmt.Sprintf("Bad line: %s", line))
		}
		lang := line[0:ldx]
		code := line[ldx+1 : cdx]
		msg := line[cdx+1 : len(line)]

		var c int
		p.Db.Model(&Locale{}).Where("lang = ? AND code = ?", lang, code).Count(&c)
		if c == 0 {
			if err = p.Db.Create(&Locale{Lang: lang, Code: code, Message: msg}).Error; err != nil {
				return err
			}
		}
	}
	return san.Err()
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
