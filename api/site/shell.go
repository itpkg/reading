package site

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/codegangsta/cli"
	"github.com/garyburd/redigo/redis"
	"github.com/itpkg/reading/api/cache"
	"github.com/itpkg/reading/api/config"
	"github.com/itpkg/reading/api/core"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"path/filepath"
	"time"
)

func writeText(tpl, htm string, model interface{}, mode os.FileMode) error {
	t, err := template.ParseFiles(path.Join("templates", tpl))
	if err != nil {
		return err
	}

	f, err := os.OpenFile(htm, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	defer f.Close()

	return t.Execute(f, model)
}

func writeAssets(dao *Dao, dist string, mode os.FileMode) error {
	ifo := make(map[string]*SiteModel)
	for _, lang := range dao.Languages() {
		ifo[lang] = &SiteModel{
			Lang:        lang,
			Title:       dao.GetSiteInfo("title", lang),
			SubTitle:    dao.GetSiteInfo("subTitle", lang),
			Keywords:    dao.GetSiteInfo("keywords", lang),
			AuthorName:  dao.GetString("authorName"),
			AuthorEmail: dao.GetString("authorEmail"),
			Description: dao.GetSiteInfo("description", lang),
			Copyright:   dao.GetSiteInfo("copyright", lang),
		}

	}

	tpl, err := getHtmlTemplate("templates")
	if err != nil {
		return err
	}

	return core.Loop(func(en core.Engine) error {
		for _, t := range en.Asserts() {
			log.Printf("Write file %s", t.Htm)
			if err := writeHtml(
				tpl,
				t.Tpl,
				dist,
				t.Htm,
				struct {
					Site    *SiteModel
					Payload interface{}
				}{Site: ifo[t.Lang], Payload: t.Payload},
				mode, true); err != nil {
				return err
			}
		}
		return nil
	})
}

func writeSeo(cfg *config.Model, dao *Dao, dist string, mode os.FileMode) error {

	googleVerify := dao.GetString("googleVerify")
	if err := writeText(
		"google_verify.tmpl",
		path.Join(dist, fmt.Sprintf("google%s.html", googleVerify)),
		googleVerify,
		mode); err != nil {
		return err
	}
	baiduVerify := dao.GetString("baiduVerify")
	if err := writeText(
		"baidu_verify.tmpl",
		path.Join(dist, fmt.Sprintf("baidu_verify_%s.html", baiduVerify)),
		baiduVerify,
		mode); err != nil {
		return err
	}
	return ioutil.WriteFile(
		path.Join(dist, "robots.txt"),
		[]byte(fmt.Sprintf("Sitemap: %s/sitemap.xml\n%s", cfg.Home(), dao.GetString("robotsTxt"))),
		mode)
}

func writeSitemap(cfg *config.Model, dist string, mode os.FileMode) error {
	root := path.Join(dist, "assets")
	home := cfg.Home()
	const header = `<?xml version="1.0" encoding="utf-8"?>`

	fd1, err := os.OpenFile(path.Join(root, "sitemap.xml"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	defer fd1.Close()
	fd1.WriteString(header)
	fd1.WriteString("\n")
	fd1.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"`)
	if err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if f.Mode().IsRegular() {
			if filepath.Ext(f.Name()) == ".html" {
				fmt.Fprintf(fd1, "<url><loc>%s%s</loc></url>", home, path[len(dist):])
			}
		}
		return nil
	}); err != nil {
		return err
	}

	fd1.WriteString("</urlset>")
	fd1.Sync()

	fd2, err := os.OpenFile(path.Join(dist, "sitemap.xml"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	fd2.WriteString(header)
	fd2.WriteString("\n")
	fd2.WriteString(`<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`)
	now := time.Now()
	for _, u := range []string{"/assets", "/api/v1"} {
		fmt.Fprintf(fd2, "<sitemap><loc>%s%s/sitemap.xml</loc><lastmod>%v</lastmod></sitemap>", home, u, now)
	}
	fd2.WriteString("</sitemapindex>")
	defer fd2.Close()
	fd2.Sync()

	return nil
}

func (p *SiteEngine) Shell() []cli.Command {
	return []cli.Command{
		{
			Name:    "assets",
			Aliases: []string{"as"},
			Usage:   "compile all the assets",
			Flags:   []cli.Flag{core.ENV},
			Action: IocAction(func(cfg *config.Model, _ *cli.Context) error {
				db, err := cfg.OpenDatabase()
				if err != nil {
					return err
				}

				const dist = "public"
				const mode = os.FileMode(0644)
				dao := Dao{Db: db}
				if err := writeAssets(&dao, dist, mode); err != nil {
					return err
				}
				if err := writeSeo(cfg, &dao, dist, mode); err != nil {
					return err
				}

				return writeSitemap(cfg, dist, mode)
			}),
		},
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
				var hnd http.Handler
				if cfg.IsProduction() {
					hnd = rt
				} else {
					//hnd = cors.Default().Handler(rt)
					hnd = cors.New(cors.Options{
						AllowedMethods: []string{"GET", "POST", "DELETE"},
						AllowedHeaders: []string{"Authorization"},
					}).Handler(rt)
				}
				return http.ListenAndServe(fmt.Sprintf(":%d", cfg.Http.Port), hnd)
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
			Aliases: []string{"ng"},
			Usage:   "generate nginx files",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "port, p",
					Value: 3000,
					Usage: "port of server to listen",
				},
				cli.StringFlag{
					Name:  "name, n",
					Value: "localhost",
					Usage: "domain name",
				},
				cli.BoolFlag{
					Name:  "ssl, s",
					Usage: "enable ssl mode",
				},
			},
			Action: core.EnvAction(func(env string, ctx *cli.Context) error {
				ssl := ctx.Bool("ssl")
				name := ctx.String("name")
				port := ctx.Int("port")

				fn := "templates/nginx/http.conf"
				if ssl {
					fn = "templates/nginx/https.conf"
				}

				t, err := template.ParseFiles(fn)
				if err != nil {
					return err
				}
				f, err := os.OpenFile("config/nginx.conf", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
				if err != nil {
					return err
				}
				defer f.Close()

				return t.Execute(f, struct {
					Name string
					Port int
				}{Name: name, Port: port})
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
