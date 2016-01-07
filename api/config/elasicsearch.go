package config

import (
	"fmt"
)

type ElasticSearch struct {
	Host  string `toml:"host"`
	Port  int    `toml:"port"`
	Index string `toml:"index"`
}

func (p *ElasticSearch) Url() string {
	return fmt.Sprintf("http://%s:%d", p.Host, p.Port)
}
