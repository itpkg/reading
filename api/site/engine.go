package site

import (
	"fmt"
	"os"
	"text/template"

	"github.com/codegangsta/cli"
	"github.com/gorilla/pat"
	"github.com/itpkg/reading/api/core"
	"github.com/itpkg/reading/api/rss"
	"github.com/itpkg/reading/api/sitemap"
	"github.com/jinzhu/gorm"
	"golang.org/x/tools/blog/atom"
)

type SiteEngine struct {
}

func (p *SiteEngine) Mount(*pat.Router) {
}

func (p *SiteEngine) Seed(*gorm.DB) error {
	//todo
	return nil

}

func (p *SiteEngine) Migrate(*gorm.DB) error {
	//todo
	return nil
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

func (p *SiteEngine) Shell() []cli.Command {
	return []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "start the web server",
			Flags:   []cli.Flag{core.ENV},
			Action: core.Action(func(env string) error {
				//				if err := Init(env); err != nil {
				//					return err
				//				}
				//				http := ioc.Get("http").(*cfg.Http)
				//
				//				if http.IsProduction() {
				//					gin.SetMode(gin.ReleaseMode)
				//				}
				//				router := gin.Default()
				//				if !http.IsProduction() {
				//					router.Static("/assets", "assets")
				//				}
				//
				//				if err := engine.Loop(func(en engine.Engine) error {
				//					en.Mount(router)
				//					return nil
				//				}); err != nil {
				//					return err
				//				}
				//
				//				return router.Run(fmt.Sprintf(":%d", http.Port))

				return nil
			}),
		},
		{
			Name:    "nginx",
			Aliases: []string{"n"},
			Usage:   "generate nginx files",
			Flags:   []cli.Flag{core.ENV},
			Action: core.Action(func(env string) error {
				cfg, e1 := core.Load(env)
				if e1 != nil {
					return e1
				}
				t, e2 := template.ParseFiles("views/nginx.conf")
				if e2 != nil {
					return e2
				}
				f, e3 := os.OpenFile("config/nginx.conf", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
				if e3 != nil {
					return e3
				}
				return t.Execute(f, cfg.Http)
			}),
		},
		{
			Name:    "database",
			Aliases: []string{"db"},
			Usage:   "database operations",
			Subcommands: []cli.Command{
				{
					Name:    "create",
					Aliases: []string{"n"},
					Usage:   "creates the database",
					Flags:   []cli.Flag{core.ENV},
					Action: core.Action(func(env string) error {
						cfg, e := core.Load(env)
						if e != nil {
							return e
						}

						c, a := cfg.Database.Execute(fmt.Sprintf("CREATE DATABASE %s WITH ENCODING='UTF8'", cfg.Database.Name))
						return core.Shell(c, a...)
					}),
				},
				{
					Name:    "console",
					Aliases: []string{"c"},
					Usage:   "start a console for the database",
					Flags:   []cli.Flag{core.ENV},
					Action: core.Action(func(env string) error {

						cfg, e := core.Load(env)
						if e != nil {
							return e
						}

						c, a := cfg.Database.Console()
						return core.Shell(c, a...)
					}),
				},
				{
					Name:    "migrate",
					Aliases: []string{"m"},
					Usage:   "migrate the database",
					Flags:   []cli.Flag{core.ENV},
					Action: core.Action(func(env string) error {

						cfg, e1 := core.Load(env)
						if e1 != nil {
							return e1
						}
						db, e2 := cfg.Database.Open()
						if e2 != nil {
							return e2
						}
						return core.Loop(func(en core.Engine) error {
							return en.Migrate(db)
						})
					}),
				},
				{
					Name:    "seed",
					Aliases: []string{"s"},
					Usage:   "load the seed data",
					Flags:   []cli.Flag{core.ENV},
					Action: core.Action(func(env string) error {

						cfg, e1 := core.Load(env)
						if e1 != nil {
							return e1
						}
						db, e2 := cfg.Database.Open()
						if e2 != nil {
							return e2
						}
						return core.Loop(func(en core.Engine) error {
							return en.Seed(db)
						})
					}),
				},
				{
					Name:    "drop",
					Aliases: []string{"d"},
					Usage:   "drops the database",
					Flags:   []cli.Flag{core.ENV},
					Action: core.Action(func(env string) error {

						cfg, err := core.Load(env)
						if err != nil {
							return err
						}
						c, a := cfg.Database.Execute(fmt.Sprintf("DROP DATABASE %s", cfg.Database.Name))
						return core.Shell(c, a...)
					}),
				},
			},
		},
	}
}

func init() {
	core.Register(&SiteEngine{})
}
