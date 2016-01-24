package auth

import (
	"net/http"
	"strings"

	"github.com/itpkg/reading/api/web"
	"github.com/julienschmidt/httprouter"
)

func (p *AuthEngine) getAdminRoles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := p.Session.Admin(r)
	if err != nil {
		p.Abort(w, err)
		return
	}
	var pms []Permission
	p.Db.Order("id DESC").Find(&pms)
	var val []interface{}
	for _, pm := range pms {
		if pm.Enable() {
			p.Db.Model(&pm).Related(&pm.User)
			p.Db.Model(&pm).Related(&pm.Role)
			val = append(val, map[string]interface{}{
				"user":  pm.User.String(),
				"role":  pm.Role.String(),
				"begin": pm.BeginS(),
				"end":   pm.EndS(),
			})
		}
	}
	p.Render.JSON(w, http.StatusOK, val)

}
func (p *AuthEngine) postAdminRoles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//todo
}
func (p *AuthEngine) getAdminSiteInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := p.Session.Admin(r)
	if err != nil {
		p.Abort(w, err)
		return
	}

	lang := p.Locale(r)
	fm := web.NewForm("siteInfo", "/admin/site/info")
	fm.Text("title", p.SiteDao.GetSiteInfo("title", lang))
	fm.Text("subTitle", p.SiteDao.GetSiteInfo("subTitle", lang))
	fm.Text("keywords", p.SiteDao.GetSiteInfo("keywords", lang))
	fm.Text("authorName", p.SiteDao.GetSiteInfo("author.name", lang))
	fm.Text("authorEmail", p.SiteDao.GetSiteInfo("author.email", lang))
	fm.TextArea("description", p.SiteDao.GetSiteInfo("description", lang))
	fm.Text("copyright", p.SiteDao.GetSiteInfo("copyright", lang))

	p.Render.JSON(w, http.StatusOK, fm)

}
func (p *AuthEngine) postAdminSiteInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//todo
}
func (p *AuthEngine) getAdminSiteSeo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := p.Session.Admin(r)
	if err != nil {
		p.Abort(w, err)
		return
	}

	fm := web.NewForm("siteSeo", "/admin/site/seo")
	fm.Text("robots", p.SiteDao.GetString("robots.txt"))
	fm.Text("googleVerify", p.SiteDao.GetString("google.verify"))
	fm.Text("baiduVerify", p.SiteDao.GetString("baidu.verify"))

	p.Render.JSON(w, http.StatusOK, fm)
}
func (p *AuthEngine) postAdminSiteSeo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//todo
}

func (p *AuthEngine) getAdminSiteSecrets(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := p.Session.Admin(r)
	if err != nil {
		p.Abort(w, err)
		return
	}
	var gcf GoogleConf
	p.SiteDao.Get("google.oauth", &gcf)

	fm := web.NewForm("siteSecrets", "/admin/site/secrets")
	fm.Text("youtubeServerKey", p.SiteDao.GetString("youtube.serverKey"))
	fm.Text("googleWebClientId", gcf.Web.ClientId)
	fm.Text("googleWebClientSecret", gcf.Web.ClientSecret)
	fm.TextArea("googleWebRedirectURLS", strings.Join(gcf.Web.RedirectURLS, "\n"))
	p.Render.JSON(w, http.StatusOK, fm)
}
func (p *AuthEngine) postAdminSiteSecrets(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//todo
}
