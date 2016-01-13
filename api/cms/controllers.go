package cms

import (
	"net/http"
	"os/exec"
	"strings"

	"github.com/itpkg/reading/api/core"
	"github.com/julienschmidt/httprouter"
)

//------------------------books------------------------------------------------
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

//------------------------articles tags comments ------------------------------
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

//------------------------videos-----------------------------------------------

func (p *CmsEngine) showChannel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var item Channel
	err := p.Db.Select([]string{"id", "cid", "title", "type", "description"}).Where("id = ?", ps.ByName("id")).First(&item).Error
	if err == nil {
		pager := core.NewPager(24)
		_, start, size := pager.Parse(r)
		var total int
		p.Db.Model(Playlist{}).Where("cid = ?", item.Cid).Count(&total)
		pager.SetTotal(total)

		var playlist []Playlist
		p.Db.Where("cid = ?", item.Cid).Offset(start).Limit(size).Find(&playlist)

		p.Render.JSON(w, http.StatusOK, map[string]interface{}{
			"pager":    pager,
			"channel":  item,
			"playlist": playlist,
		})

	} else {
		p.Abort(w, err)
	}
}
func (p *CmsEngine) listChannel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pager := core.NewPager(24)
	_, start, size := pager.Parse(r)
	var total int
	p.Db.Model(Channel{}).Count(&total)
	pager.SetTotal(total)

	var items []Channel
	p.Db.Select([]string{"id", "cid", "title", "type", "description"}).Offset(start).Limit(size).Find(&items)
	p.Pager(p.Render, w, pager, items)
}
func (p *CmsEngine) showPlaylist(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var item Playlist
	err := p.Db.Select([]string{"id", "pid", "title", "type", "description"}).Where("id = ?", ps.ByName("id")).First(&item).Error
	if err == nil {
		pager := core.NewPager(24)
		_, start, size := pager.Parse(r)
		var total int
		p.Db.Model(Video{}).Where("pid = ?", item.Pid).Count(&total)
		pager.SetTotal(total)

		var videos []Video
		p.Db.Where("pid = ?", item.Pid).Offset(start).Limit(size).Find(&videos)

		p.Render.JSON(w, http.StatusOK, map[string]interface{}{
			"pager":    pager,
			"playlist": item,
			"videos":   videos,
		})
	} else {
		p.Abort(w, err)
	}
}
func (p *CmsEngine) listPlaylist(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pager := core.NewPager(24)
	_, start, size := pager.Parse(r)
	var total int
	p.Db.Model(Playlist{}).Count(&total)
	pager.SetTotal(total)

	var items []Playlist
	p.Db.Select([]string{"id", "pid", "title", "type", "description"}).Offset(start).Limit(size).Find(&items)
	p.Pager(p.Render, w, pager, items)
}
func (p *CmsEngine) showVideo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var item Video
	err := p.Db.Select([]string{"id", "vid", "title", "type", "description"}).Where("id = ?", ps.ByName("id")).First(&item).Error
	if err == nil {
		p.Render.JSON(w, http.StatusOK, item)
	} else {
		p.Abort(w, err)
	}
}
func (p *CmsEngine) listVideo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pager := core.NewPager(24)
	_, start, size := pager.Parse(r)
	var total int
	p.Db.Model(Video{}).Count(&total)
	pager.SetTotal(total)

	var items []Video
	p.Db.Select([]string{"id", "vid", "title", "type", "description"}).Offset(start).Limit(size).Find(&items)
	p.Pager(p.Render, w, pager, items)
}

func (p *CmsEngine) dict(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if buf, err := exec.Command("sdcv", "--data-dir", "dict", r.URL.Query().Get("keyword")).Output(); err == nil {
		p.Html(w, strings.Replace(string(buf), "\n", "<br>", -1))
	} else {
		p.Abort(w, err)
	}

}

func (p *CmsEngine) Mount(rt core.Router) {
	rt.GET("/videos/channels/:id", p.showChannel)
	rt.GET("/videos/channels", p.listChannel)
	rt.GET("/videos/playlist/:id", p.showPlaylist)
	rt.GET("/videos/playlist", p.listPlaylist)
	rt.GET("/videos/items/:id", p.showVideo)
	rt.GET("/videos/items", p.listVideo)

	rt.GET("/cms/books/:id", p.showBook)
	rt.GET("/cms/books", p.listBook)

	rt.GET("/cms/articles/:id", p.showArticle)
	rt.GET("/cms/articles", p.listArticle)
	rt.GET("/cms/tags/:id", p.showTag)
	rt.GET("/cms/tags", p.listTag)

	rt.POST("/dict", p.dict)
}
