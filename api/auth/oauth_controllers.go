package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (p *AuthEngine) postSignIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()

	flag := r.Form.Get("type")
	code := r.Form.Get("code")
	var token string
	var err error
	switch flag {
	case "google":
		token, err = p.googleCallback(code)
	default:
		err = errors.New(fmt.Sprintf("Unsupported oauth %s", flag))
	}

	if err == nil {
		p.Render.JSON(w, http.StatusOK, token)
	} else {
		p.Abort(w, err)
	}
}

func (p *AuthEngine) googleCallback(code string) (string, error) {
	cfg, err := p.googleConf()
	if err != nil {
		return "", err
	}
	g := NewGoogle(cfg)

	buf, err := g.Token(code)
	if err != nil {
		return "", err
	}
	info := make(map[string]interface{})
	if err := json.Unmarshal(buf, &info); err != nil {
		return "", err
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
		return "", err
	}
	p.Dao.Log(user.ID, "sign in")

	return p.Token.New(map[string]interface{}{
		"id":      user.Uid,
		"name":    user.Name,
		"isAdmin": p.Dao.Is(user.ID, "admin"),
		//	"logo": user.Logo,
	}, 7*24*60)
}

func (p *AuthEngine) getSignIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c, e := p.googleConf()
	if e != nil {
		p.Abort(w, e)
		return
	}

	g := NewGoogle(c)

	p.Render.JSON(w, http.StatusOK, map[string]string{
		"google": g.Url(),
	})
}

func (p *AuthEngine) googleConf() (*GoogleConf, error) {
	cfg := GoogleConf{}
	err := p.SiteDao.Get("google.oauth", &cfg)

	return &cfg, err
}
