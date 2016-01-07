package auth

import (
	"encoding/json"
	"os"
	"time"

	"github.com/codegangsta/cli"
	"github.com/itpkg/reading/api/config"
	"github.com/itpkg/reading/api/core"
	"github.com/itpkg/reading/api/rss"
	"github.com/itpkg/reading/api/site"
	"github.com/itpkg/reading/api/sitemap"
	"github.com/itpkg/reading/api/token"
	"github.com/jinzhu/gorm"
	"github.com/unrolled/render"
	"golang.org/x/tools/blog/atom"
)

type AuthEngine struct {
	core.Controller

	SiteDao *site.Dao      `inject:""`
	Db      *gorm.DB       `inject:""`
	Dao     *Dao           `inject:""`
	Cfg     *config.Model  `inject:""`
	Token   token.Provider `inject:""`
	Render  *render.Render `inject:""`
}

func (p *AuthEngine) loadGoogleOauthConf() (*GoogleConf, error) {
	cf, err := os.Open("config/google.oauth.json")
	if err != nil {
		return nil, err
	}

	de := json.NewDecoder(cf)
	cfg := GoogleConf{}
	if err := de.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
func (p *AuthEngine) initRoles() error {

	db := p.Db
	var count int
	db.Model(Role{}).Count(&count)
	if count == 0 {
		dur := 24 * 365 * 10 * time.Hour

		var rootR *Role
		var adminR *Role
		var err error
		if rootR, err = p.Dao.GetRole("root", "-", 0); err != nil {
			return err
		}
		if err = p.Dao.Apply(rootR.ID, 1, dur); err != nil {
			return err
		}
		if adminR, err = p.Dao.GetRole("admin", "-", 0); err != nil {
			return err
		}
		if err = p.Dao.Apply(adminR.ID, 1, dur); err != nil {
			return err
		}
	}
	return nil
}
func (p *AuthEngine) Seed() error {
	if gcf, err := p.loadGoogleOauthConf(); err == nil {
		if err := p.SiteDao.Set("google.oauth", gcf, true); err != nil {
			return err
		}
	} else {
		return err
	}

	return p.initRoles()
}

func (p *AuthEngine) Migrate() {
	db := p.Db
	db.AutoMigrate(&User{}, &Role{}, &Permission{}, &Log{})
	db.Model(&User{}).AddUniqueIndex("idx_user_provider_type_id", "provider_type", "provider_id")
	db.Model(&Role{}).AddUniqueIndex("idx_roles_name_resource_type_id", "name", "resource_type", "resource_id")
	db.Model(&Permission{}).AddUniqueIndex("idx_permissions_user_role", "user_id", "role_id")

}

func (p *AuthEngine) Sitemap() sitemap.Handler {
	return func() []*sitemap.Url {
		return []*sitemap.Url{}
	}
}
func (p *AuthEngine) Rss() rss.Handler {
	return func(lang string) []*atom.Entry {
		return []*atom.Entry{}
	}
}

func (p *AuthEngine) Shell() []cli.Command {
	return []cli.Command{}
}

func init() {
	core.Register(&AuthEngine{})
}
