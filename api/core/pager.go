package core

import (
	"net/http"
	"strconv"
)

type Pager struct {
	Total int
	Size  int
}

func (p *Pager) FromAndSize(r *http.Request) (int, int) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil || size <= 0 || size >= p.Size {
		page = 20
	}

	count := p.Total / p.Size
	if p.Total%p.Size != 0 {
		count++
	}

	if page*size > p.Total {
		page = count
	}

	return (page - 1) * size, size
}
