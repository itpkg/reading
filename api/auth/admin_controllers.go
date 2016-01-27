package auth

import (
	"net/http"
	"strings"

	"github.com/itpkg/reading/api/site"
	"github.com/itpkg/reading/api/web"
	"github.com/julienschmidt/httprouter"
)

func (p *AuthEngine) locales(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := p.Session.Admin(r)
	if err != nil {
		p.Abort(w, err)
		return
	}
	var items []site.Locale
	p.Db.Select([]string{"id", "code", "message"}).Where("lang = ?", p.Locale(r)).Order("code DESC").Find(&items)
	p.Render.JSON(w, http.StatusOK, items)
}

func (p *AuthEngine) saveLocale(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := p.Session.Admin(r)
	if err != nil {
		p.Abort(w, err)
		return
	}
	r.ParseForm()

	id := r.FormValue("id")
	lang := p.Locale(r)
	code := r.FormValue("code")
	message := r.FormValue("message")

	if id == "" {
		err = p.Db.Create(&site.Locale{Lang: lang, Code: code, Message: message}).Error
	} else {
		var lol site.Locale
		p.Db.Where("id = ?", id).First(&lol)
		lol.Code = code
		lol.Message = message
		err = p.Db.Save(&lol).Error
	}
	res := web.NewResponse(true, nil)
	res.Check(err)

	p.Render.JSON(w, http.StatusOK, res)

}

func (p *AuthEngine) destroyLocale(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, err := p.Session.Admin(r)
	if err != nil {
		p.Abort(w, err)
		return
	}

	p.Db.Where("id = ?", ps.ByName("id")).Delete(site.Locale{})
	p.Render.JSON(w, http.StatusOK, web.NewResponse(true, nil))
}

func (p *AuthEngine) saveNotice(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := p.Session.Admin(r)
	if err != nil {
		p.Abort(w, err)
		return
	}
	r.ParseForm()
	id := r.FormValue("id")
	content := r.FormValue("content")
	lang := p.Locale(r)
	if id == "" {
		err = p.Db.Create(&site.Notice{Content: content, Lang: lang}).Error
	} else {
		var n site.Notice
		p.Db.Where("id = ?", id).First(&n)
		n.Content = content
		err = p.Db.Save(&n).Error
	}

	res := web.NewResponse(true, nil)
	res.Check(err)
	p.Render.JSON(w, http.StatusOK, res)
}
func (p *AuthEngine) destroyNotice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, err := p.Session.Admin(r)
	if err != nil {
		p.Abort(w, err)
		return
	}

	p.Db.Where("id = ?", ps.ByName("id")).Delete(site.Notice{})
	p.Render.JSON(w, http.StatusOK, web.NewResponse(true, nil))
}

func (p *AuthEngine) getAdminUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := p.Session.Admin(r)
	if err != nil {
		p.Abort(w, err)
		return
	}
	var users []User
	p.Db.Select([]string{"id", "name", "email", "last_sign_in"}).Order("sign_in_count DESC").Find(&users)
	var val []interface{}
	for _, u := range users {
		var pms []map[string]string
		p.Db.Model(&u).Related(&u.Permissions)
		for _, pm := range u.Permissions {
			p.Db.Model(pm).Related(&pm.Role)
			pms = append(pms, map[string]string{
				"role":  pm.Role.String(),
				"begin": pm.BeginS(),
				"end":   pm.EndS(),
			})
		}
		val = append(val, map[string]interface{}{
			"label":       u.String(),
			"lastSignIn":  u.LastSignIn,
			"permissions": pms,
		})
	}
	p.Render.JSON(w, http.StatusOK, val)

}
func (p *AuthEngine) getAdminSiteInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := p.Session.Admin(r)
	if err != nil {
		p.Abort(w, err)
		return
	}

	lang := p.Locale(r)
	fm := web.NewForm("siteInfo", "/admin/site/info")
	for _, k := range []string{
		"title",
		"subTitle",
		"keywords",
		"copyright",
		"authorName", "authorEmail",
	} {
		fm.Text(k, p.SiteDao.GetSiteInfo(k, lang))
	}

	fm.TextArea("description", p.SiteDao.GetSiteInfo("description", lang))

	p.Render.JSON(w, http.StatusOK, fm)
}
func (p *AuthEngine) postAdminSiteInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := p.Session.Admin(r)
	if err != nil {
		p.Abort(w, err)
		return
	}

	lang := p.Locale(r)
	r.ParseForm()
	for _, k := range []string{
		"title",
		"subTitle",
		"keywords",
		"description",
		"copyright",
		"authorName", "authorEmail",
	} {
		p.SiteDao.SetSiteInfo(k, lang, r.FormValue(k), false)
	}
	p.Render.JSON(w, http.StatusOK, web.NewResponse(true, nil))
}

func (p *AuthEngine) getAdminSiteSeo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := p.Session.Admin(r)
	if err != nil {
		p.Abort(w, err)
		return
	}

	ifo := make(map[string]string)
	for _, k := range []string{"googleVerify", "baiduVerify", "robotsTxt"} {
		ifo[k] = p.SiteDao.GetString(k)
	}
	p.Render.JSON(w, http.StatusOK, ifo)

	//	fm := web.NewForm("siteSeo", "/admin/site/seo")
	//	fm.Text("googleVerify", p.SiteDao.GetString("googleVerify"))
	//	fm.Text("baiduVerify", p.SiteDao.GetString("baiduVerify"))
	//	fm.TextArea("robotsTxt", p.SiteDao.GetString("robotsTxt"))
	//
	//	p.Render.JSON(w, http.StatusOK, fm)
}
func (p *AuthEngine) postAdminSiteSeo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := p.Session.Admin(r)
	if err != nil {
		p.Abort(w, err)
		return
	}

	r.ParseForm()
	for _, k := range []string{
		"robotsTxt",
		"googleVerify",
		"baiduVerify",
	} {
		p.SiteDao.Set(k, r.FormValue(k), false)
	}
	p.Render.JSON(w, http.StatusOK, web.NewResponse(true, nil))
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
	fm.Text("youtubeServerKey", p.SiteDao.GetString("youtubeServerKey"))
	fm.Text("googleWebClientId", gcf.Web.ClientId)
	fm.Text("googleWebClientSecret", gcf.Web.ClientSecret)
	fm.TextArea("googleWebRedirectURLS", strings.Join(gcf.Web.RedirectURLS, "\n"))
	p.Render.JSON(w, http.StatusOK, fm)
}
func (p *AuthEngine) postAdminSiteSecrets(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := p.Session.Admin(r)
	if err != nil {
		p.Abort(w, err)
		return
	}
	r.ParseForm()
	gcf := GoogleConf{
		Web: GoogleWeb{
			ClientId:     r.FormValue("googleWebClientId"),
			ClientSecret: r.FormValue("googleWebClientSecret"),
			RedirectURLS: strings.Split(r.FormValue("googleWebRedirectURLS"), "\n"),
		},
	}
	p.SiteDao.Set("youtubeServerKey", r.FormValue("youtubeServerKey"), true)
	p.SiteDao.Set("google.oauth", &gcf, true)
	p.Render.JSON(w, http.StatusOK, web.NewResponse(true, nil))
}

func (p *AuthEngine) getAdminSiteTop(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := p.Session.Admin(r)
	if err != nil {
		p.Abort(w, err)
		return
	}

	fm := web.NewForm("siteTop", "/admin/site/top")
	fm.TextArea("content", p.SiteDao.GetString("topNavBar"))
	p.Render.JSON(w, http.StatusOK, fm)
}
func (p *AuthEngine) postAdminSiteTop(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := p.Session.Admin(r)
	if err != nil {
		p.Abort(w, err)
		return
	}

	r.ParseForm()
	p.SiteDao.Set("topNavBar", r.FormValue("content"), false)

	p.Render.JSON(w, http.StatusOK, web.NewResponse(true, nil))
}
