package site

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (p *SiteEngine) sitemap(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	//todo
}

func (p *SiteEngine) rss(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	//todo
}

func (p *SiteEngine) info(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	lang := p.Locale(r)
	p.Render.JSON(w, http.StatusOK, map[string]interface{}{
		"title":       p.Dao.GetSiteInfo("title", lang),
		"subTitle":    p.Dao.GetSiteInfo("sub_title", lang),
		"keywords":    p.Dao.GetSiteInfo("keywords", lang),
		"description": p.Dao.GetSiteInfo("description", lang),
		"author": map[string]string{
			"email": p.Dao.GetSiteInfo("author.email", lang),
			"name":  p.Dao.GetSiteInfo("author.name", lang),
		},
		"copyright": p.Dao.GetSiteInfo("copyright", lang),
	})
}

func (p *SiteEngine) robots(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	var tx string
	p.Dao.Get("robots.txt", &tx)
	p.Render.Text(w, http.StatusOK, tx)
}

func (p *SiteEngine) baidu(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	c := p.Dao.GetSiteInfo("baidu.site.verify", "")
	if fmt.Sprintf("%s.html", c) == ps.ByName("id") {
		p.Render.HTML(w, http.StatusOK, "baidu_verify", c)
	} else {
		p.NotFound(w)
	}
}

func (p *SiteEngine) google(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	c := p.Dao.GetSiteInfo("google.site.verify", "")
	if fmt.Sprintf("%s.html", c) == ps.ByName("id") {
		p.Render.HTML(w, http.StatusOK, "google_verify", c)
	} else {
		p.NotFound(w)
	}
}
