package core

import (
	"net/http"
	"strconv"
)

func NewPager(max int) *Pager {
	return &Pager{Size: max}
}

type Pager struct {
	Size  int `json:"size"`
	Total int `json:"total"`
	Count int `json:"count"`
	Page  int `json:"page"`
}

func (p *Pager) Parse(r *http.Request) (int, int, int) {

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil || size <= 0 || size > p.Size {
		size = p.Size
	}

	p.Page = page
	p.Size = size
	return page, (page - 1) * size, size
}

func (p *Pager) To() map[string]int {
	return map[string]int{
		"total": p.Count,
		"size":  p.Size,
		"page":  p.Page,
		"count": p.Count,
	}
}

func (p *Pager) SetTotal(total int) int {

	count := total / p.Size
	if total%p.Size != 0 {
		count++
	}

	p.Total = total
	p.Count = count
	return count
}
