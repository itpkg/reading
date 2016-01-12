package cms

import (
	"log"
	"time"

	"github.com/codegangsta/cli"
	"github.com/itpkg/reading/api/config"
	"github.com/itpkg/reading/api/core"
	"github.com/itpkg/reading/api/epub"
	"github.com/itpkg/reading/api/site"
	"github.com/jinzhu/gorm"
)

func (p *CmsEngine) Shell() []cli.Command {
	return []cli.Command{
		{
			Name:  "cms",
			Usage: "cms module tasks",
			Subcommands: []cli.Command{
				{
					Name:    "books",
					Aliases: []string{"b"},
					Usage:   "import books",
					Flags:   []cli.Flag{core.ENV},
					Action: config.DatabaseAction(func(db *gorm.DB, _ *cli.Context) error {
						root := "tmp/books"
						return epub.Walk(root, func(name string) error {
							log.Printf("find file %s", name)
							it, err := epub.Open(name)
							//							log.Printf("%+v", it.Container)
							opf := it.Container.RootFiles[0].Opf
							index := opf.Index()
							cover := opf.Cover()
							book := Book{
								Model: core.Model{
									UpdatedAt: time.Now(),
								},
								Name:       name[len(root)+1 : len(name)-5],
								Title:      opf.Metadata.Creator,
								Author:     opf.Metadata.Creator,
								Language:   opf.Metadata.Language,
								Subject:    opf.Metadata.Subject,
								Publisher:  opf.Metadata.Publisher,
								Date:       opf.Metadata.Date,
								Identifier: opf.Metadata.Identifier,
								Type:       it.MimeType,
								IndexHref:  index.Href,
								IndexType:  index.MediaType,
								CoverHref:  cover.Href,
								CoverType:  cover.MediaType,
							}
							//log.Printf("%+v", book)
							var c int
							db.Model(&Book{}).Where("name = ?", book.Name).Count(&c)
							if c == 0 {
								db.Create(&book)
							} else {
								db.Model(Book{Name: book.Name}).Where("name = ?", book.Name).Updates(book)
							}

							return err
						})
					}),
				},
				{
					Name:    "videos",
					Aliases: []string{"v"},
					Usage:   "import videos",
					Flags:   []cli.Flag{core.ENV},
					Action: site.IocAction(func(*config.Model, *cli.Context) error {
						//todo
						return nil
					}),
				},
			},
		},
	}
}
