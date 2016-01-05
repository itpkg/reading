package rss

import (
	"encoding/xml"
	"io"
	"time"

	"github.com/pborman/uuid"
	"golang.org/x/tools/blog/atom"
)

type Handler func(string) []*atom.Entry

func Xml(wrt io.Writer, lang, title, home, username, email string, handlers ...Handler) error {
	feed := atom.Feed{
		Title: title,
		ID:    uuid.New(),
		Link: []atom.Link{
			{
				Href: home,
			},
		},
		Updated: atom.Time(time.Now()),
		Author: &atom.Person{
			Name:  username,
			Email: email,
		},
		Entry: make([]*atom.Entry, 0),
	}
	for _, fn := range handlers {
		items := fn(lang)
		feed.Entry = append(feed.Entry, items...)
	}

	if _, err := wrt.Write([]byte(xml.Header)); err != nil {
		return err
	}
	en := xml.NewEncoder(wrt)
	//en.Indent("", "  ")
	return en.Encode(feed)
}
