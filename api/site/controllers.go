package site

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/itpkg/reading/api/core"
	"github.com/itpkg/reading/api/rss"
	"github.com/itpkg/reading/api/sitemap"
	"github.com/julienschmidt/httprouter"
)

func (p *SiteEngine) sitemap(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	p.Cache.Page(w, r, "application/xml", 24*60, func() ([]byte, error) {
		var buf bytes.Buffer
		hds := make([]sitemap.Handler, 0)
		core.Loop(func(en core.Engine) error {
			hds = append(hds, en.Sitemap())
			return nil
		})
		err := sitemap.Xml(&buf, hds...)
		return buf.Bytes(), err
	})
}

func (p *SiteEngine) rss(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	lang := p.Locale(r)
	p.Cache.Page(w, r, "application/xml", 6*60, func() ([]byte, error) {
		var buf bytes.Buffer
		hds := make([]rss.Handler, 0)
		core.Loop(func(en core.Engine) error {
			hds = append(hds, en.Rss())
			return nil
		})
		err := rss.Xml(
			&buf,
			lang,
			p.Dao.GetSiteInfo("title", lang),
			p.Cfg.Home(),
			p.Dao.GetSiteInfo("author.name", ""),
			p.Dao.GetSiteInfo("author.email", ""),
			hds...,
		)
		return buf.Bytes(), err
	})
}

func (p *SiteEngine) info(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	lang := p.Locale(r)
	p.Cache.Page(w, r, "application/json", 6*60, func() ([]byte, error) {
		ifo := map[string]interface{}{
			"title":       p.Dao.GetSiteInfo("title", lang),
			"subTitle":    p.Dao.GetSiteInfo("sub_title", lang),
			"keywords":    p.Dao.GetSiteInfo("keywords", lang),
			"description": p.Dao.GetSiteInfo("description", lang),
			"author": map[string]string{
				"email": p.Dao.GetSiteInfo("author.email", lang),
				"name":  p.Dao.GetSiteInfo("author.name", lang),
			},
			"copyright": p.Dao.GetSiteInfo("copyright", lang),
		}
		return core.ToJson(ifo)
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
