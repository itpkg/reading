package config

type Storage struct {
	Type  string            `toml:"type"`
	Extra map[string]string `toml:"extra"`
}
