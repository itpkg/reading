package config

import (
	"fmt"

	"github.com/itpkg/reading/api/core"
)

func Load(env string) (*Model, error) {
	cfg := Model{Env: env}
	if err := core.FromToml(fmt.Sprintf("config/%s.toml", env), &cfg); err == nil {
		cfg.Secrets, err = core.FromBase64(cfg.SecretsS)
		return &cfg, err
	} else {
		return nil, err
	}
}
