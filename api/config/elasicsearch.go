package config

type ElasticSearch struct {
	Host  string `toml:"host"`
	Port  int    `toml:"port"`
	Index string `toml:"index"`
}
