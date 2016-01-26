package cms

import (
	"net/http"

	"github.com/itpkg/reading/api/core"
	"github.com/julienschmidt/httprouter"
)

func (p *CmsEngine) showTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var item Tag
	err := p.Db.First(&item, ps.ByName("id")).Error
	if err != nil {
		p.Abort(w, err)
	}

	if err = p.Db.Model(&item).Association("Articles").Error; err == nil {
		p.Render.JSON(w, http.StatusOK, item)
	} else {
		p.Abort(w, err)
	}
}

func (p *CmsEngine) listTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pager := core.NewPager(60)
	_, start, size := pager.Parse(r)
	var total int
	p.Db.Model(Tag{}).Count(&total)
	pager.SetTotal(total)

	var items []Tag
	p.Db.Select([]string{"id", "name"}).Offset(start).Limit(size).Find(&items)
	p.Pager(p.Render, w, pager, items)
}
