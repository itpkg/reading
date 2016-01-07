package assets

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/itpkg/reading/api/config"
	"github.com/itpkg/reading/api/core"
	"gopkg.in/olivere/elastic.v3"
)

func (p *AssetsEngine) Shell() []cli.Command {
	return []cli.Command{

		{
			Name:    "assets",
			Aliases: []string{"as"},
			Usage:   "assets manage",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Usage:   "list asset items",
					Flags: []cli.Flag{
						core.ENV,
						cli.IntFlag{
							Name:  "from, f",
							Value: 0,
							Usage: "from of pager",
						},
						cli.IntFlag{
							Name:  "size, s",
							Value: 20,
							Usage: "size of pager",
						},
					},
					Action: config.ElasticAction(func(con *elastic.Client, cfg *config.Model, ctx *cli.Context) error {

						dao := Dao{Client: con, Cfg: cfg}
						fmt.Println("ID TITLE")
						_, items, err := dao.List("", ctx.Int("from"), ctx.Int("size"))
						if err != nil {
							return err
						}
						for _, it := range items {
							fmt.Printf("%s %s\n", it.Id(), it.Title)
						}
						return nil
					}),
				},
			},
		},
	}
}
