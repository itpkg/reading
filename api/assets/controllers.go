package assets

import (
	"net/http"

	"github.com/itpkg/reading/api/core"
	"github.com/julienschmidt/httprouter"
)

const ROOT = "/assets/"

func (p *AssetsEngine) showBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//todo
}
func (p *AssetsEngine) listBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//todo
	var books []Book
	p.Db.Find(&books)
	p.Render.JSON(w, http.StatusOK, books)
}
func (p *AssetsEngine) showBlog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
func (p *AssetsEngine) listBlog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var items []Asset
	p.Db.Select([]string{"id", "title"}).Where("type = ? OR type = ?", "markdown", "latex").Find(&items)
	p.Render.JSON(w, http.StatusOK, items)
}
func (p *AssetsEngine) showChannel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//todo
}
func (p *AssetsEngine) listChannel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//todo
}
func (p *AssetsEngine) showPlaylist(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//todo
}
func (p *AssetsEngine) listPlaylist(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//todo
}
func (p *AssetsEngine) showVideo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//todo
}
func (p *AssetsEngine) listVideo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//todo
}

func (p *AssetsEngine) Mount(rt core.Router) {
	rt.GET("/channels/:cid", p.showChannel)
	rt.GET("/channels", p.listChannel)
	rt.GET("/playlist/:cid", p.showPlaylist)
	rt.GET("/playlist", p.listPlaylist)
	rt.GET("/videos/:cid", p.showVideo)
	rt.GET("/videos", p.listVideo)

	rt.GET("/books/*url", p.showBook)
	rt.GET("/books", p.listBook)

	rt.GET("/blog/*url", p.showBlog)
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
