package core

import (
	"crypto/cipher"
	"fmt"

	"github.com/itpkg/reading/api/config"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
)

func Load(env string) (*config.Model, error) {
	cfg := config.Model{Env: env}
	if err := FromToml(fmt.Sprintf("config/%s.toml", env), &cfg); err == nil {
		cfg.Secrets, err = FromBase64(cfg.SecretsS)
		return &cfg, err
	} else {
		return nil, err
	}
}

func Init(env string) error {
	var err error

	//--------config
	var cfg *config.Model
	if cfg, err = Load(env); err != nil {
		return err
	}

	//-----------------aes
	var cip cipher.Block
	if cip, err = NewAesCipher(cfg.Secrets[60:92]); err != nil {
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
	if db, err = cfg.Database.Open(); err != nil {
		return err
	}
	db.LogMode(!cfg.IsProduction())

	//------------
	if err = In(
		cfg,
		db,
		cfg.Redis.Open(),
		logger,
	); err != nil {
		return err
	}
	if err = Use(map[string]interface{}{
		"aes.cipher": cip,
	}); err != nil {
		return err
	}
	if err = Loop(func(en Engine) error {
		return In(en)
	}); err != nil {
		return err
	}

	return Check()
}
