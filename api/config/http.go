package config

type Http struct {
	Domain string `toml:"domain"`
	Port   int    `toml:"port"`
	Ssl    bool   `toml:"ssl"`
}
