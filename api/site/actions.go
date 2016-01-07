package site

import (
	"github.com/codegangsta/cli"
	"github.com/itpkg/reading/api/config"
)

func IocAction(act func(*config.Model, *cli.Context) error) func(c *cli.Context) {
	return config.ConfigAction(func(cfg *config.Model, ctx *cli.Context) error {
		if err := Init(cfg.Env); err == nil {
			return act(cfg, ctx)
		} else {
			return err
		}
	})
}
