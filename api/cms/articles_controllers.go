package cms

import (
	"net/http"

	"github.com/itpkg/reading/api/core"
	"github.com/julienschmidt/httprouter"
)

func (p *CmsEngine) showArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var item Article
	err := p.Db.First(&item, ps.ByName("id")).Error
	if err != nil {
		p.Abort(w, err)
	}

	if err = p.Db.Model(&item).Association("Tags").Error; err == nil {
		p.Render.JSON(w, http.StatusOK, item)
	} else {
		p.Abort(w, err)
	}
}

func (p *CmsEngine) listArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pager := core.NewPager(60)
	_, start, size := pager.Parse(r)
	var total int
	p.Db.Model(Article{}).Count(&total)
	pager.SetTotal(total)

	var items []Article
	p.Db.Select([]string{"id", "summary", "title"}).Offset(start).Limit(size).Find(&items)
	p.Pager(p.Render, w, pager, items)
}
