package cms

import (
	"net/http"

	"github.com/itpkg/reading/api/core"
	"github.com/julienschmidt/httprouter"
)

func (p *CmsEngine) showBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var item Book
	err := p.Db.Where("id = ?", ps.ByName("id")).First(&item).Error
	if err == nil {
		p.Render.JSON(w, http.StatusOK, item)
	} else {
		p.Abort(w, err)
	}
}
func (p *CmsEngine) listBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pager := core.NewPager(24)
	_, start, size := pager.Parse(r)
	var total int
	p.Db.Model(Book{}).Count(&total)
	pager.SetTotal(total)

	var items []Book
	p.Db.Offset(start).Limit(size).Find(&items)
	p.Pager(p.Render, w, pager, items)
}
