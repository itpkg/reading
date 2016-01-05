package config

import (
	"fmt"
)

type Model struct {
	Env           string         `toml:"-"`
	SecretsS      string         `toml:"secret"`
	Secrets       []byte         `toml:"-"`
	Http          *Http          `toml:"http"`
	Database      *Database      `toml:"database"`
	Redis         *Redis         `toml:"redis"`
	ElasticSearch *ElasticSearch `toml:"elastic_search"`
}

func (p *Model) IsProduction() bool {
	return p.Env == "production"
}

func (p *Model) Home() string {
	if p.IsProduction() {
		if p.Http.Ssl {
			return fmt.Sprintf("https://www.%s", p.Http.Domain)
		} else {
			return fmt.Sprintf("http://www.%s", p.Http.Domain)
		}
	} else {
		return fmt.Sprintf("http://localhost:%s", p.Http.Port)
	}
}
