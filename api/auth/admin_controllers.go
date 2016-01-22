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
	//todo
}
func (p *AuthEngine) postAdminSiteInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//todo
}
func (p *AuthEngine) getAdminSiteSeo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//todo
}
func (p *AuthEngine) postAdminSiteSeo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//todo
}
