package cms

import (
	"net/http"

	"github.com/itpkg/reading/api/core"
	"github.com/julienschmidt/httprouter"
)

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
