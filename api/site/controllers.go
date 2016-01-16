package site

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/itpkg/reading/api/core"
	"github.com/itpkg/reading/api/rss"
	"github.com/itpkg/reading/api/sitemap"
	"github.com/julienschmidt/httprouter"
)

func (p *SiteEngine) sitemap(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	p.Cache.Page(w, r, "application/xml", 24*60, func() ([]byte, error) {
		var buf bytes.Buffer
		var hds []sitemap.Handler
		core.Loop(func(en core.Engine) error {
			hds = append(hds, en.Sitemap())
			return nil
		})
		err := sitemap.Xml(&buf, hds...)
		return buf.Bytes(), err
	})
}

func (p *SiteEngine) rss(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	p.Cache.Page(w, r, "application/xml", 6*60, func() ([]byte, error) {
		lang := p.Locale(r)
		var buf bytes.Buffer
		var hds []rss.Handler
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

func (p *SiteEngine) info(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	p.Cache.Page(w, r, "application/json", 6*60, func() ([]byte, error) {
		lang := p.Locale(r)
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

func (p *SiteEngine) locales(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	p.Cache.Page(w, r, "application/json", 6*60, func() ([]byte, error) {
		lang := ps.ByName("lang")

		var items []Locale
		p.Db.Select("code, message").Where("code LIKE ? AND lang = ?", "web.%", lang).Order("code").Find(&items)

		rt := make(map[string]interface{})
		for _, item := range items {
			codes := strings.Split(item.Code[4:], ".")
			tmp := rt
			for i, c := range codes {
				if i+1 == len(codes) {
					tmp[c] = item.Message
				} else {
					if tmp[c] == nil {
						tmp[c] = make(map[string]interface{})
					}
					tmp = tmp[c].(map[string]interface{})
				}
			}
		}
		return core.ToJson(rt)
	})

}

func (p *SiteEngine) Mount(rt core.Router) {
	//just for i18next
	rt.GET("/locales/:lang", p.locales)

	rt.GET("/site/info", p.info)
	rt.GET("/baidu_verify_:id", p.baidu)
	rt.GET("/google:id", p.google)
	rt.GET("/rss.atom", p.rss)
	rt.GET("/sitemap.xml", p.sitemap)
	rt.GET("/robots.txt", p.robots)
}
