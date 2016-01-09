package assets

import (
	"net/http"

	"github.com/itpkg/reading/api/core"
	"github.com/julienschmidt/httprouter"
)

func (p *AssetsEngine) showBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var item Book
	err := p.Db.Select([]string{"id", "url", "title", "author"}).Where("id = ?", ps.ByName("id")).First(&item).Error
	if err == nil {
		p.Render.JSON(w, http.StatusOK, item)
	} else {
		p.Abort(w, err)
	}
}
func (p *AssetsEngine) listBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pager := core.NewPager(24)
	_, start, size := pager.Parse(r)
	var total int
	p.Db.Model(Book{}).Count(&total)
	pager.SetTotal(total)

	var items []Book
	p.Db.Select([]string{"id", "url", "title", "author"}).Offset(start).Limit(size).Find(&items)
	p.Pager(p.Render, w, pager, items)
}
func (p *AssetsEngine) showBlog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var item Page
	err := p.Db.Select([]string{"id", "title"}).Where("id = ?", ps.ByName("id")).First(&item).Error
	if err == nil {
		p.Render.JSON(w, http.StatusOK, item)
	} else {
		p.Abort(w, err)
	}

}
func (p *AssetsEngine) listBlog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pager := core.NewPager(60)
	_, start, size := pager.Parse(r)
	var total int
	p.Db.Model(Page{}).Count(&total)
	pager.SetTotal(total)

	var items []Page
	p.Db.Select([]string{"id", "title"}).Where("type = ? OR type = ?", "markdown", "latex").Offset(start).Limit(size).Find(&items)
	p.Pager(p.Render, w, pager, items)
}
func (p *AssetsEngine) showChannel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
func (p *AssetsEngine) listChannel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pager := core.NewPager(24)
	_, start, size := pager.Parse(r)
	var total int
	p.Db.Model(Channel{}).Count(&total)
	pager.SetTotal(total)

	var items []Channel
	p.Db.Select([]string{"id", "cid", "title", "type", "description"}).Offset(start).Limit(size).Find(&items)
	p.Pager(p.Render, w, pager, items)
}
func (p *AssetsEngine) showPlaylist(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
func (p *AssetsEngine) listPlaylist(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pager := core.NewPager(24)
	_, start, size := pager.Parse(r)
	var total int
	p.Db.Model(Playlist{}).Count(&total)
	pager.SetTotal(total)

	var items []Playlist
	p.Db.Select([]string{"id", "pid", "title", "type", "description"}).Offset(start).Limit(size).Find(&items)
	p.Pager(p.Render, w, pager, items)
}
func (p *AssetsEngine) showVideo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var item Video
	err := p.Db.Select([]string{"id", "vid", "title", "type", "description"}).Where("id = ?", ps.ByName("id")).First(&item).Error
	if err == nil {
		p.Render.JSON(w, http.StatusOK, item)
	} else {
		p.Abort(w, err)
	}
}
func (p *AssetsEngine) listVideo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pager := core.NewPager(24)
	_, start, size := pager.Parse(r)
	var total int
	p.Db.Model(Video{}).Count(&total)
	pager.SetTotal(total)

	var items []Video
	p.Db.Select([]string{"id", "vid", "title", "type", "description"}).Offset(start).Limit(size).Find(&items)
	p.Pager(p.Render, w, pager, items)
}

func (p *AssetsEngine) Mount(rt core.Router) {
	rt.GET("/channels/:id", p.showChannel)
	rt.GET("/channels", p.listChannel)
	rt.GET("/playlist/:id", p.showPlaylist)
	rt.GET("/playlist", p.listPlaylist)
	rt.GET("/videos/:id", p.showVideo)
	rt.GET("/videos", p.listVideo)

	rt.GET("/books/:id", p.showBook)
	rt.GET("/books", p.listBook)

	rt.GET("/blog/:id", p.showBlog)
	rt.GET("/blog", p.listBlog)
}

/*
func (p *AssetsEngine) index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		pager := core.NewPager(20)
		_, from, size := pager.Parse(r)
		total, items, err := p.Dao.List("text/html", from, size)
		if err == nil {
			pager.SetTotal(int(total))
			p.Render.JSON(w, http.StatusOK, map[string]interface{}{
				"pager": pager.To(),
				"items": items,
			})

		} else {
			p.Abort(w, err)
		}
}

func (p *AssetsEngine) show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		it := Item{Title: ps.ByName("name")}
		if item, err := p.Dao.Get(it.Id()); err == nil {
			p.Render.JSON(w, http.StatusOK, item)
		} else {
			p.Abort(w, err)
		}
}

*/
