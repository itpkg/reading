package site

import (
	"github.com/itpkg/reading/api/core"
)

type Setting struct {
	core.Model
	Key  string `sql:"not null;unique"`
	Val  []byte `sql:"not null"`
	Flag bool   `sql:"not null;default:false"`
}

type Locale struct {
	ID      uint   `gorm:"primary_key"`
	Code    string `sql:"not null;index:idx_locales_code"`
	Lang    string `sql:"not null;default:'en-US';type:varchar(5);index:idx_locales_lang"`
	Message string `sql:"not null;type:text"`
}