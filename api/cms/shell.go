package cms

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"fmt"
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
						return fetch_epub_books(db, "tmp/books")
					}),
				},
				{
					Name:    "videos",
					Aliases: []string{"v"},
					Usage:   "import videos",
					Flags:   []cli.Flag{core.ENV},
					Action: config.ConfigAction(func(cfg *config.Model, _ *cli.Context) error {
						db, err := cfg.OpenDatabase()
						if err != nil {
							return err
						}
						cip, err := cfg.AesCipher()
						if err != nil {
							return err
						}
						dao := site.Dao{Db: db, Aes: &core.Aes{Cip: cip}}
						var key string
						if err := dao.Get("youtube.key", &key); err != nil {
							return err
						}
						var users []User
						db.Select([]string{"uid", "type"}).Find(&users)
						for _, user := range users {
							log.Printf("Find user %s", user)
							switch user.Type {
							case "youtube":
								if err := fetch_youtube_user(db, key, user.Uid); err != nil {
									return err
								}
							default:
								log.Println("Ingnore")
							}
						}

						return nil
					}),
				},
			},
		},
	}
}

func http_get_json(url string, obj interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	log.Printf("GET %s\n%s", url, res.Body)

	dec := json.NewDecoder(res.Body)
	err = dec.Decode(obj)
	if err != nil {
		log.Println(res.Body)
	}
	return err
}

//--------------------YOUTUBE-------------
//todo 分页
func fetch_youtube_user(db *gorm.DB, key, uid string) error {
	var user YTUser
	err := http_get_json(
		fmt.Sprintf(
			"https://www.googleapis.com/youtube/v3/channels?part=snippet&maxResults=50&forUsername=%s&key=%s",
			uid,
			key),
		&user)
	if err != nil {
		return err
	}

	for _, item := range user.Items {
		ch := Channel{
			Model: core.Model{
				UpdatedAt: time.Now(),
			},
			Type:        "youtube",
			Uid:         uid,
			Cid:         item.Id,
			Title:       item.Snippet.Title,
			Description: item.Snippet.Description,
		}
		log.Printf("get channel %s", ch.Title)
		var c int
		db.Model(&Channel{}).Where("cid = ? AND type = ?", ch.Cid, "youtube").Count(&c)
		if c == 0 {
			err = db.Save(&ch).Error
		} else {
			err = db.Model(Channel{}).Where("cid = ? AND type = ?", ch.Cid, "youtube").Updates(ch).Error
		}
		if err != nil {
			return err
		}
		if err := fetch_youtube_channel(db, key, ch.Cid); err != nil {
			return err
		}
	}
	return nil
}

func fetch_youtube_channel(db *gorm.DB, key, cid string) error {
	var ch YTChannel
	err := http_get_json(
		fmt.Sprintf(
			"https://www.googleapis.com/youtube/v3/playlists?part=snippet&channelId=%s&maxResults=50&key=%s",
			cid,
			key),
		&ch,
	)
	if err != nil {
		return nil
	}
	for _, item := range ch.Items {
		pl := Playlist{
			Model: core.Model{
				UpdatedAt: time.Now(),
			},
			Type:        "youtube",
			Cid:         cid,
			Pid:         item.Id,
			Title:       item.Snippet.Title,
			Description: item.Snippet.Description,
		}
		log.Printf("get playlist %s", pl.Title)
		var c int
		db.Model(&Playlist{}).Where("pid = ? AND type = ?", pl.Pid, "youtube").Count(&c)
		if c == 0 {
			err = db.Save(&pl).Error
		} else {
			err = db.Model(Playlist{}).Where("pid = ? AND type = ?", pl.Pid, "youtube").Updates(pl).Error
		}
		if err != nil {
			return err
		}
		if err = fetch_youtube_playlist(db, key, pl.Pid); err != nil {
			return err
		}
	}
	return nil
}

func fetch_youtube_playlist(db *gorm.DB, key, pid string) error {
	var pl YTPlaylist
	err := http_get_json(
		fmt.Sprintf(
			"https://www.googleapis.com/youtube/v3/playlistItems?part=snippet&maxResults=50&playlistId=%s&key=%s",
			pid,
			key),
		&pl,
	)
	if err != nil {
		return nil
	}
	for _, item := range pl.Items {
		v := Video{
			Model: core.Model{
				UpdatedAt: time.Now(),
			},
			Type:        "youtube",
			Pid:         pid,
			Vid:         item.Id,
			Title:       item.Snippet.Title,
			Description: item.Snippet.Description,
		}
		log.Printf("get video %s", v.Title)
		var c int
		db.Model(&Video{}).Where("vid = ? AND type = ?", v.Vid, "youtube").Count(&c)
		if c == 0 {
			err = db.Save(&v).Error
		} else {
			err = db.Model(Video{}).Where("vid = ? AND type = ?", v.Vid, "youtube").Updates(v).Error
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func fetch_epub_books(db *gorm.DB, root string) error {
	return epub.Walk(root, func(name string) error {
		log.Printf("find file %s", name)
		it, err := epub.Open(name)
		if err != nil {
			return err
		}
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
			err = db.Create(&book).Error
		} else {
			err = db.Model(Book{}).Where("name = ?", book.Name).Updates(book).Error
		}

		return err
	})

}
