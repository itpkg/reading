package site

import (
	"github.com/itpkg/reading/api/cache"
	"github.com/itpkg/reading/api/config"
	"github.com/itpkg/reading/api/core"
	"github.com/itpkg/reading/api/token"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/unrolled/render"
)

func Init(env string) error {

	//--------config

	cfg, err := config.Load(env)
	if err != nil {
		return err
	}

	//---------storage
	stg, err := cfg.OpenStorage()
	if err != nil {
		return err
	}

	//-----------------aes
	cip, err := cfg.AesCipher()
	if err != nil {
		return err
	}

	//-----------------logger
	var logger = logging.MustGetLogger("reading")
	if cfg.IsProduction() {
		if bkd, err := logging.NewSyslogBackend("itpkg"); err == nil {
			logging.SetBackend(bkd)
		} else {
			return err
		}
	}

	//--------------database
	var db *gorm.DB
	if db, err = cfg.OpenDatabase(); err != nil {
		return err
	}
	//--------------cache
	var cp cache.Provider
	cp = &cache.RedisProvider{}
	//--------------Elastic search
	//	esc, err := cfg.OpenElastic()
	//	if err != nil {
	//		return err
	//	}

	//------------
	if err = core.In(
		cfg,
		db,
		cfg.OpenRedis(),
		cp,
		//		esc,
		&token.RedisProvider{},
		stg,
		render.New(render.Options{
			IsDevelopment: !cfg.IsProduction(),
		}),
		logger,
	); err != nil {
		return err
	}
	if err = core.Use(map[string]interface{}{
		"aes.cipher": cip,
	}); err != nil {
		return err
	}
	if err = core.Loop(func(en core.Engine) error {
		return core.In(en)
	}); err != nil {
		return err
	}

	return core.Check()
}
