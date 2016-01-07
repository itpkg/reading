package auth

import (
	"encoding/json"
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
	info := make(map[string]interface{})
	if err := json.Unmarshal(buf, &info); err != nil {
		p.Abort(w, err)
		return
	}

	user, err := p.Dao.SaveUser(
		"google",
		info["id"].(string),
		info["email"].(string),
		info["name"].(string),
		info["link"].(string),
		info["picture"].(string),
	)
	if err != nil {
		p.Abort(w, err)
		return
	}

	tkn, err := p.Token.New(map[string]interface{}{
		"id":   user.Uid,
		"name": user.Name,
		"logo": user.Logo,
	}, 7*24*60)
	p.Render.HTML(w, http.StatusOK, "sign_in", tkn)
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
