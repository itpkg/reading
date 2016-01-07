package site

import (
	"github.com/codegangsta/cli"
	"github.com/itpkg/reading/api/core"
)

func IocAction(act func(*cli.Context) error) func(c *cli.Context) {
	return core.EnvAction(func(env string, ctx *cli.Context) error {
		if err := Init(env); err == nil {
			return act(ctx)
		} else {
			return err
		}
	})
}
