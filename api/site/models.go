package site

import (
	"github.com/itpkg/reading/api/core"
)

type Setting struct {
	core.Model
	Key  string `sql:"not null;unique_index"`
	Val  []byte `sql:"not null"`
	Flag bool   `sql:"not null;default:false"`
}

type Locale struct {
	core.Model
	Code    string `sql:"not null;index:idx_locales_code" json:"code"`
	Lang    string `sql:"not null;default:'en-US';type:char(5);index:idx_locales_lang" json:"lang"`
	Message string `sql:"not null;type:text" json:"message"`
}

type Notice struct {
	core.Model
	Lang    string `sql:"not null;type:char(5);index:idx_notices_lang" json:"lang"`
	Content string `sql:"not null;type:text" json:"content"`
}
