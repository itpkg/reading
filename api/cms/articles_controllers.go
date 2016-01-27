package cms

import (
	"net/http"

	"github.com/itpkg/reading/api/core"
	"github.com/itpkg/reading/api/web"
	"github.com/julienschmidt/httprouter"
)

func (p *CmsEngine) showArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var item Article
	err := p.Db.Where("aid = ?", ps.ByName("aid")).First(&item).Error
	if err != nil {
		p.Abort(w, err)
		return
	}

	if err = p.Db.Model(&item).Association("Tags").Find(&item.Tags).Error; err == nil {
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
func (p *CmsEngine) saveArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user, err := p.Session.User(r)
	if err != nil {
		p.Abort(w, err)
		return
	}
	r.ParseForm()
	aid := r.FormValue("aid")
	id := r.FormValue("id")
	title := r.FormValue("title")
	body := r.FormValue("body")
	summary := r.FormValue("summary")

	if id == "" {
		err = p.Db.Create(&Article{
			Aid:     aid,
			UserID:  user.ID,
			Title:   title,
			Summary: summary,
			Body:    body,
		}).Error
	} else {
		var a Article
		p.Db.Where("id = ?", id).First(&a)
		if a.UserID != user.ID && !p.AuthDao.Is(user.ID, "admin") {
			p.Forbidden(w)
			return
		}
		a.Aid = aid
		a.Title = title
		a.Summary = summary
		a.Body = body
		err = p.Db.Save(&a).Error
	}
	rsp := web.NewResponse(true, nil)
	rsp.Check(err)
	p.Render.JSON(w, http.StatusOK, rsp)
}
