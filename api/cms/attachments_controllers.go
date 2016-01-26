package cms

import (
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/itpkg/reading/api/core"
	"github.com/itpkg/reading/api/web"
	"github.com/julienschmidt/httprouter"
)

func (p *CmsEngine) storeAttachment(user uint, fh *multipart.FileHeader) error {

	fd, err := fh.Open()
	defer fd.Close()
	if err != nil {
		return err
	}

	url, size, err := p.Storage.Store(fh.Filename, fd)
	if err != nil {
		return err
	}

	hds := fh.Header["Content-Type"]
	if len(hds) == 0 {
		return errors.New("bad header")
	}
	p.Db.Create(&Attachment{
		UserID: user,
		Title:  fh.Filename,
		Type:   hds[0],
		Url:    url,
		Size:   size,
	})

	return nil
}

func (p *CmsEngine) saveAttachment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	user, err := p.Session.User(r)
	if err != nil {
		p.Abort(w, err)
		return
	}
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		p.Abort(w, err)
		return
	}

	for _, fhs := range r.MultipartForm.File {
		for _, fh := range fhs {
			if err := p.storeAttachment(user.ID, fh); err != nil {
				p.Abort(w, err)
				return
			}
		}
	}

	p.Render.JSON(w, http.StatusOK, web.NewResponse(true, nil))

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
		p.Db.Where("id = ?", ps.ByName("id")).Delete(Attachment{})
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

	user, err := p.Session.User(r)
	if err != nil {
		p.Abort(w, err)
		return
	}
	db := p.Db.Select([]string{"url", "title", "id", "created_at"})
	if !p.AuthDao.Is(user.ID, "admin") {
		db = db.Where("user_id = ?", user.ID)
	}
	var items []Attachment
	db.Offset(start).Order("id DESC").Limit(size).Find(&items)
	p.Pager(p.Render, w, pager, items)
}
