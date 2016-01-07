package config

type Redis struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
	Db   int    `toml:"db"`
}
