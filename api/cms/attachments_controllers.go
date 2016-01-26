package cms

import (
	"net/http"

	"github.com/itpkg/reading/api/core"
	"github.com/itpkg/reading/api/site"
	"github.com/itpkg/reading/api/web"
	"github.com/julienschmidt/httprouter"
)

func (p *CmsEngine) saveAttachment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//	user, err := p.Session.User(r)
	//	if err != nil {
	//		p.Abort(w, err)
	//		return
	//	}

}

func (p *CmsEngine) removeAttachment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user, err := p.Session.User(r)
	if err != nil {
		p.Abort(w, err)
		return
	}
	var a Attachment

	if err := p.Db.Where("id = ?", ps.ByName("id")).First(&a).Error; err != nil {
		p.Abort(w, err)
		return
	}

	if p.AuthDao.Is(user.ID, "admin") || user.ID == a.UserID {
		p.Db.Where("id = ?", ps.ByName("id")).Delete(site.Locale{})
		p.Render.JSON(w, http.StatusOK, web.NewResponse(true, nil))
	} else {
		p.Forbidden(w)
	}

}

func (p *CmsEngine) listAttachment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pager := core.NewPager(60)
	_, start, size := pager.Parse(r)
	var total int
	p.Db.Model(Attachment{}).Count(&total)
	pager.SetTotal(total)

	var items []Attachment
	p.Db.Offset(start).Limit(size).Find(&items)
	p.Pager(p.Render, w, pager, items)
}
