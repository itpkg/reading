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
	err := FromToml(fmt.Sprintf("config/%s.toml", env), &cfg)
	return &cfg, err
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
	//--------------database
	var db *gorm.DB
	if db, err = cfg.Database.Open(); err != nil {
		return err
	}
	if !cfg.IsProduction() {
		db.LogMode(true)
	}

	//-----------------logger
	var logger = logging.MustGetLogger("portal")
	if cfg.IsProduction() {
		if bkd, err := logging.NewSyslogBackend("itpkg"); err == nil {
			logging.SetBackend(bkd)
		} else {
			return err
		}
	}
	//------------

	if err = In(
		cfg,
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

	return Check()
}
