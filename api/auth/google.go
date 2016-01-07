package auth

/**
https://developers.google.com/identity/protocols/OAuth2WebServer
https://developers.google.com/identity/protocols/googlescopes
*/

import (
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleConf struct {
	Web GoogleWeb `json:"web"`
}

type GoogleWeb struct {
	ClientId     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	RedirectURLS []string `json:"redirect_uris"`
}

type Google struct {
	cfg *oauth2.Config
}

func (p *Google) Token(req *http.Request) ([]byte, error) {
	tok, err := p.cfg.Exchange(oauth2.NoContext, req.URL.Query().Get("code"))
	if err != nil {
		return nil, err
	}
	client := p.cfg.Client(oauth2.NoContext, tok)
	res, err := client.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var buf []byte
	buf, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (p *Google) Url() string {
	return p.cfg.AuthCodeURL("state")
}

//===============================================
func NewGoogle(cfg *GoogleConf) *Google {
	return &Google{
		cfg: &oauth2.Config{
			ClientID:     cfg.Web.ClientId,
			ClientSecret: cfg.Web.ClientSecret,
			RedirectURL:  cfg.Web.RedirectURLS[0],
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.profile",
				"https://www.googleapis.com/auth/userinfo.email",
			},
			Endpoint: google.Endpoint,
		},
	}
}
