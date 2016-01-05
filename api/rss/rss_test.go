package rss_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/itpkg/reading/api/rss"
	"golang.org/x/tools/blog/atom"
)

func items(lang string) []*atom.Entry {
	items := make([]*atom.Entry, 0)

	for i := 0; i < 10; i++ {
		items = append(items, &atom.Entry{
			ID: fmt.Sprintf("article_%d", i),
			Link: []atom.Link{
				{
					Href: fmt.Sprintf("http://localhost/%d", i),
				},
			},
			Title: fmt.Sprintf("Title %d", i),
			Summary: &atom.Text{
				Body: fmt.Sprintf("Summary %d", i),
			},
			Content: &atom.Text{
				Body: fmt.Sprintf("Content %d", i),
			},
		})

	}
	return items
}

func TestAtom(t *testing.T) {
	rss.Xml(os.Stdout, "zh-CN", "ttttttt", "http://localhost", "whoami", "whoami@gmail.com", items)

}
