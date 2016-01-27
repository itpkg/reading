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
func (p *CmsEngine) listArticleBySelf(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user, err := p.Session.User(r)
	if err != nil {
		p.Abort(w, err)
		return
	}

	pager := core.NewPager(60)
	_, start, size := pager.Parse(r)
	var total int
	count := p.Db.Model(Article{})
	all := p.Db.Select([]string{"aid", "summary", "title"}).Order("updated_at DESC").Offset(start).Limit(size)
	if !p.AuthDao.Is(user.ID, "admin") {
		count = count.Where("user_id = ?", user.ID)
		all = all.Where("user_id = ?", user.ID)
	}
	count.Count(&total)
	pager.SetTotal(total)

	var items []Article
	all.Find(&items)
	p.Pager(p.Render, w, pager, items)
}
func (p *CmsEngine) listArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pager := core.NewPager(60)
	_, start, size := pager.Parse(r)
	var total int
	p.Db.Model(Article{}).Count(&total)
	pager.SetTotal(total)

	var items []Article
	p.Db.Select([]string{"aid", "summary", "title"}).Order("updated_at DESC").Offset(start).Limit(size).Find(&items)
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

	lang := p.Locale(r)
	var tags []Tag
	for _, tn := range r.Form["tags[]"] {
		var t Tag
		if p.Db.Where("name = ?", tn).First(&t).RecordNotFound() {
			t.Name = tn
			p.Db.Create(&t)
		}
		tags = append(tags, t)
	}

	if id == "" {
		err = p.Db.Create(&Article{
			Aid:     aid,
			UserID:  user.ID,
			Title:   title,
			Summary: summary,
			Body:    body,
			Lang:    lang,
			Tags:    tags,
		}).Error
	} else {
		var a Article
		p.Db.Where("id = ?", id).First(&a)
		if a.UserID != user.ID && !p.AuthDao.Is(user.ID, "admin") {
			p.Forbidden(w)
			return
		}
		p.Db.Model(&a).Association("Tags").Clear()
		a.Aid = aid
		a.Title = title
		a.Summary = summary
		a.Body = body
		a.Tags = tags
		err = p.Db.Save(&a).Error
	}
	rsp := web.NewResponse(true, nil)
	rsp.Check(err)
	p.Render.JSON(w, http.StatusOK, rsp)
}
