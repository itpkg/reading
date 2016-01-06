package site

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/codegangsta/cli"
	"github.com/itpkg/reading/api/cache"
	"github.com/itpkg/reading/api/core"
	"github.com/julienschmidt/httprouter"
	"log"
)

func (p *SiteEngine) Shell() []cli.Command {
	return []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "start the web server",
			Flags:   []cli.Flag{core.ENV},
			Action: core.Action(func(env string) error {
				if err := core.Init(env); err != nil {
					return err
				}
				cfg, err := core.Load(env)
				if err != nil {
					return err
				}
				rt := httprouter.New()
				core.Loop(func(en core.Engine) error {
					en.Mount(rt)
					return nil
				})
				return http.ListenAndServe(fmt.Sprintf(":%d", cfg.Http.Port), rt)
			}),
		},
		{
			Name:    "routers",
			Aliases: []string{"ro"},
			Usage:   "print out all defined routes in match order, with names",
			Flags:   []cli.Flag{core.ENV},
			Action: func(c *cli.Context) {
				rt := router{routes: make([]*route, 0)}
				core.Loop(func(en core.Engine) error {
					en.Mount(&rt)
					return nil
				})
				for _, r := range rt.routes {
					name := core.FuncName(r.handle)
					fmt.Printf("%s\t%s\t%v\n", r.method, r.path, name[0:strings.Index(name, ")")+1])
				}
			},
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
			Name:    "cache",
			Aliases: []string{"c"},
			Usage:   "cache operations",
			Subcommands: []cli.Command{
				{
					Name:    "remove",
					Aliases: []string{"r"},
					Usage:   "remove cache item by key",
					Flags: []cli.Flag{
						core.ENV,
						cli.StringFlag{
							Name:  "key, k",
							Value: "",
							Usage: "key of cache item",
						},
					},
					Action: func(c *cli.Context) {
						cfg, err := core.Load(c.String("environment"))
						if err != nil {
							log.Fatalln(err)
						}
						cp := cache.RedisProvider{Redis: cfg.Redis.Open()}
						if err = cp.Del(c.String("key")); err != nil {
							log.Fatalln(err)
						}
					},
				},
				{
					Name:    "list",
					Aliases: []string{"l"},
					Usage:   "list all cache keys",
					Flags:   []cli.Flag{core.ENV},
					Action: core.Action(func(env string) error {
						cfg, err := core.Load(env)
						if err != nil {
							return err
						}
						cp := cache.RedisProvider{Redis: cfg.Redis.Open()}
						keys, err := cp.Status()
						if err == nil {
							fmt.Println("TTL\tKEY")
							for k, t := range keys {
								fmt.Printf("%d\t%s\n", t, k)
							}
						}
						return err
					}),
				},
				{
					Name:    "clear",
					Aliases: []string{"c"},
					Usage:   "clear all cache items",
					Flags:   []cli.Flag{core.ENV},
					Action: core.Action(func(env string) error {
						cfg, err := core.Load(env)
						if err != nil {
							return err
						}
						cp := cache.RedisProvider{Redis: cfg.Redis.Open()}
						return cp.Clear()
					}),
				},
			},
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
						if err := core.Init(env); err != nil {
							return err
						}
						return core.Loop(func(en core.Engine) error {
							en.Migrate()
							return nil
						})
					}),
				},
				{
					Name:    "seed",
					Aliases: []string{"s"},
					Usage:   "load the seed data",
					Flags:   []cli.Flag{core.ENV},
					Action: core.Action(func(env string) error {
						if err := core.Init(env); err != nil {
							return err
						}
						return core.Loop(func(en core.Engine) error {
							return en.Seed()
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
