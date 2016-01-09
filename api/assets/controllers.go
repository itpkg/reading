package assets

import (
	"github.com/itpkg/reading/api/core"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (p *AssetsEngine) index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//	pager := core.NewPager(20)
	//	_, from, size := pager.Parse(r)
	//	total, items, err := p.Dao.List("text/html", from, size)
	//	if err == nil {
	//		pager.SetTotal(int(total))
	//		p.Render.JSON(w, http.StatusOK, map[string]interface{}{
	//			"pager": pager.To(),
	//			"items": items,
	//		})
	//
	//	} else {
	//		p.Abort(w, err)
	//	}
}

func (p *AssetsEngine) show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//	it := Item{Title: ps.ByName("name")}
	//	if item, err := p.Dao.Get(it.Id()); err == nil {
	//		p.Render.JSON(w, http.StatusOK, item)
	//	} else {
	//		p.Abort(w, err)
	//	}
}

func (p *AssetsEngine) Mount(rt core.Router) {
	rt.GET("/assets/*name", p.show)
	rt.GET("/assets", p.index)
}
