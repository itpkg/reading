package assets

import (
	"github.com/itpkg/reading/api/core"
)

type Item struct {
	Title string `json:"title"`
	Type  string `json:"type"`
	Body  string `json:"body"`
}

func (p *Item) Id() string {
	return core.Md5([]byte(p.Title))
}
