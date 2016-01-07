package config

import (
	"github.com/codegangsta/cli"
	"github.com/garyburd/redigo/redis"
	"github.com/itpkg/reading/api/core"
	"github.com/jinzhu/gorm"
	"gopkg.in/olivere/elastic.v3"
)

func ConfigAction(act func(*Model, *cli.Context) error) func(c *cli.Context) {
	return core.EnvAction(func(env string, ctx *cli.Context) error {
		if cfg, err := Load(env); err == nil {
			return act(cfg, ctx)
		} else {
			return err
		}
	})
}

func DatabaseAction(act func(*gorm.DB, *cli.Context) error) func(c *cli.Context) {
	return ConfigAction(func(cfg *Model, ctx *cli.Context) error {
		if db, err := cfg.OpenDatabase(); err == nil {
			return act(db, ctx)
		} else {
			return err
		}
	})
}

func RedisAction(act func(*redis.Pool, *cli.Context) error) func(c *cli.Context) {
	return ConfigAction(func(cfg *Model, ctx *cli.Context) error {
		return act(cfg.OpenRedis(), ctx)
	})
}

func ElasticAction(act func(*elastic.Client, *Model, *cli.Context) error) func(c *cli.Context) {
	return ConfigAction(func(cfg *Model, ctx *cli.Context) error {
		if con, err := cfg.OpenElastic(); err == nil {
			return act(con, cfg, ctx)
		} else {
			return err
		}
	})
}
