package auth

import (
	"net/http"

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
	p.Render.JSON(w, http.StatusOK, map[string]interface{}{
		"title":       p.SiteDao.GetSiteInfo("title", lang),
		"subTitle":    p.SiteDao.GetSiteInfo("subTitle", lang),
		"copyright":   p.SiteDao.GetSiteInfo("copyright", lang),
		"keywords":    p.SiteDao.GetSiteInfo("keywords", lang),
		"description": p.SiteDao.GetSiteInfo("description", lang),
		"author": map[string]string{
			"user":  p.SiteDao.GetSiteInfo("author.user", lang),
			"email": p.SiteDao.GetSiteInfo("author.email", lang),
		},
	})

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

	p.Render.JSON(w, http.StatusOK, map[string]interface{}{
		"robots": p.SiteDao.GetString("robots.txt"),
		"google": map[string]string{
			"verify": p.SiteDao.GetString("google.verify"),
		},
		"baidu": map[string]string{
			"baidu": p.SiteDao.GetString("baidu.verify"),
		},
	})
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

	p.Render.JSON(w, http.StatusOK, map[string]interface{}{
		"youtube": map[string]string{
			"key": p.SiteDao.GetString("youtube.key"),
		},
		"google": map[string]interface{}{
			"oauth": gcf,
		},
	})
}
func (p *AuthEngine) postAdminSiteSecrets(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//todo
}
