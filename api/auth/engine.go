package auth

import (
	"fmt"
	"time"

	"github.com/codegangsta/cli"
	"github.com/itpkg/reading/api/config"
	"github.com/itpkg/reading/api/core"
	"github.com/itpkg/reading/api/rss"
	"github.com/itpkg/reading/api/sitemap"
	"github.com/jinzhu/gorm"
	"golang.org/x/tools/blog/atom"
)

type AuthEngine struct {
	Db  *gorm.DB      `inject:""`
	Dao *Dao          `inject:""`
	Cfg *config.Model `inject:""`
}

func (p *AuthEngine) Seed() error {

	var count int
	p.Db.Model(User{}).Count(&count)
	if count == 0 {
		var root *User
		var adminR *Role
		var rootR *Role
		var err error
		if root, err = p.Dao.NewEmailUser("root", fmt.Sprintf("root@%s", p.Cfg.Http.Domain), "changeme"); err != nil {
			return err
		}

		dur := 24 * 365 * 10 * time.Hour

		if err = p.Dao.ConfirmUser(root.ID); err != nil {
			return err
		}
		if rootR, err = p.Dao.NewRole("root", "-", 0); err != nil {
			return err
		}
		if err = p.Dao.Apply(rootR.ID, root.ID, dur); err != nil {
			return err
		}
		if adminR, err = p.Dao.NewRole("admin", "-", 0); err != nil {
			return err
		}
		if err = p.Dao.Apply(adminR.ID, root.ID, dur); err != nil {
			return err
		}

	}
	return nil

}
func (p *AuthEngine) Migrate() {
	db := p.Db
	db.AutoMigrate(&User{}, &Contact{}, &Role{}, &Permission{}, &Log{})
	db.Model(&User{}).AddUniqueIndex("idx_user_provider_type_id", "provider_type", "provider_id")
	db.Model(&Role{}).AddUniqueIndex("idx_roles_name_resource_type_id", "name", "resource_type", "resource_id")
	db.Model(&Permission{}).AddUniqueIndex("idx_permissions_user_role", "user_id", "role_id")

}

func (p *AuthEngine) Sitemap() sitemap.Handler {
	return func() []*sitemap.Url {
		return []*sitemap.Url{} //todo
	}
}
func (p *AuthEngine) Rss() rss.Handler {
	return func(lang string) []*atom.Entry {
		return []*atom.Entry{} //todo
	}
}

func (p *AuthEngine) Shell() []cli.Command {
	return []cli.Command{}
}

func init() {
	core.Register(&AuthEngine{})
}
