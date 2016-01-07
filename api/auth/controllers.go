package auth

import (
	"net/http"

	"github.com/itpkg/reading/api/core"
	"github.com/julienschmidt/httprouter"
)

func (p *AuthEngine) googleCallback(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cfg, e := p.getGoogleConf()
	if e != nil {
		p.Abort(w, e)
		return
	}
	g := NewGoogle(cfg)

	buf, e2 := g.Token(r)
	if e2 != nil {
		p.Abort(w, e2)
		return
	}
	w.Write(buf)
}

func (p *AuthEngine) googleSignIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c, e := p.getGoogleConf()
	if e != nil {
		p.Abort(w, e)
		return
	}
	g := NewGoogle(c)

	http.Redirect(w, r, g.Url(), http.StatusFound)
}

func (p *AuthEngine) getGoogleConf() (*GoogleConf, error) {
	cfg := GoogleConf{}
	err := p.SiteDao.Get("google.oauth", &cfg)

	return &cfg, err
}

func (p *AuthEngine) Mount(rt core.Router) {
	rt.GET("/oauth/google/sign_in", p.googleSignIn)
	rt.GET("/oauth/google/callback", p.googleCallback)
}
