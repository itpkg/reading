package config

import (
	"fmt"

	"github.com/itpkg/reading/api/core"
)

type Http struct {
	Env     string `toml:"-"`
	Domain  string `toml:"domain"`
	Port    int    `toml:"port"`
	Secrets string `toml:"secret"`
	Ssl     bool   `toml:"ssl"`
}

func (p *Http) Home() string {
	if p.IsProduction() {
		if p.Ssl {
			return fmt.Sprintf("https://www.%s", p.Domain)
		} else {
			return fmt.Sprintf("http://www.%s", p.Domain)
		}
	} else {
		return fmt.Sprintf("http://localhost:%s", p.Port)
	}
}

func (p *Http) Key(begin, end int) ([]byte, error) {
	if b, e := core.FromBase64(p.Secrets); e == nil {
		return b[begin:end], nil
	} else {
		return nil, e
	}

}

func (p *Http) IsProduction() bool {
	return p.Env == "production"
}
