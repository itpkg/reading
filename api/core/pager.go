package core

import (
	"net/http"
	"strconv"
)

func NewPager(max int) *Pager {
	return &Pager{size: max}
}

type Pager struct {
	size  int
	total int
	count int
	page  int
}

func (p *Pager) Parse(r *http.Request) (int, int, int) {

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil || size <= 0 || size > p.size {
		size = p.size
	}

	p.page = page
	p.size = size
	return page, (page - 1) * size, size
}

func (p *Pager) To() map[string]int {
	return map[string]int{
		"total": p.count,
		"size":  p.size,
		"page":  p.page,
		"count": p.count,
	}
}

func (p *Pager) SetTotal(total int) int {

	count := total / p.size
	if total%p.size != 0 {
		count++
	}

	p.total = total
	p.count = count
	return count
}
