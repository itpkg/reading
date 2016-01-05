package config

type Model struct {
	Http          *Http          `toml:"http"`
	Database      *Database      `toml:"database"`
	Redis         *Redis         `toml:"redis"`
	ElasticSearch *ElasticSearch `toml:"elastic_search"`
}
