package site

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/codegangsta/cli"
	"github.com/garyburd/redigo/redis"
	"github.com/itpkg/reading/api/cache"
	"github.com/itpkg/reading/api/config"
	"github.com/itpkg/reading/api/core"
	"github.com/julienschmidt/httprouter"
)

func (p *SiteEngine) Shell() []cli.Command {
	return []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "start the web server",
			Flags:   []cli.Flag{core.ENV},
			Action: IocAction(func(cfg *config.Model, _ *cli.Context) error {
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
			Action: config.ConfigAction(func(cfg *config.Model, ctx *cli.Context) error {
				t, e2 := template.ParseFiles("templates/nginx.conf")
				if e2 != nil {
					return e2
				}
				f, e3 := os.OpenFile("config/nginx.conf", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
				defer f.Close()
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
					Action: config.RedisAction(func(rep *redis.Pool, ctx *cli.Context) error {
						key := ctx.String("key")
						if key == "" {
							return errors.New("need a key")
						}
						cp := cache.RedisProvider{Redis: rep}
						return cp.Del(key)
					}),
				},
				{
					Name:    "list",
					Aliases: []string{"l"},
					Usage:   "list all cache keys",
					Flags:   []cli.Flag{core.ENV},
					Action: config.RedisAction(func(rep *redis.Pool, ctx *cli.Context) error {
						cp := cache.RedisProvider{Redis: rep}
						keys, err := cp.Status()
						if err != nil {
							return err
						}
						fmt.Println("TTL\tKEY")
						for k, t := range keys {
							fmt.Printf("%d\t%s\n", t, k)
						}

						return nil
					}),
				},
				{
					Name:    "clear",
					Aliases: []string{"c"},
					Usage:   "clear all cache items",
					Flags:   []cli.Flag{core.ENV},
					Action: config.RedisAction(func(rep *redis.Pool, ctx *cli.Context) error {
						cp := cache.RedisProvider{Redis: rep}
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
					Action: config.ConfigAction(func(cfg *config.Model, ctx *cli.Context) error {
						c, a := cfg.Database.Execute(fmt.Sprintf("CREATE DATABASE %s WITH ENCODING='UTF8'", cfg.Database.Name))
						return core.Shell(c, a...)
					}),
				},
				{
					Name:    "console",
					Aliases: []string{"c"},
					Usage:   "start a console for the database",
					Flags:   []cli.Flag{core.ENV},
					Action: config.ConfigAction(func(cfg *config.Model, ctx *cli.Context) error {
						c, a := cfg.Database.Console()
						return core.Shell(c, a...)
					}),
				},
				{
					Name:    "migrate",
					Aliases: []string{"m"},
					Usage:   "migrate the database",
					Flags:   []cli.Flag{core.ENV},
					Action: IocAction(func(*config.Model, *cli.Context) error {
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
					Action: IocAction(func(*config.Model, *cli.Context) error {
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
					Action: config.ConfigAction(func(cfg *config.Model, ctx *cli.Context) error {
						c, a := cfg.Database.Execute(fmt.Sprintf("DROP DATABASE %s", cfg.Database.Name))
						return core.Shell(c, a...)
					}),
				},
			},
		},
	}
}
