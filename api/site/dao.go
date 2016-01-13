package site

import (
	"fmt"
	"strings"

	"github.com/itpkg/reading/api/core"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	Db  *gorm.DB  `inject:""`
	Aes *core.Aes `inject:""`
}

func (p *Dao) SetSiteInfo(key, lang string, val interface{}, flag bool) error {
	return p.Set(p.site_key(key, lang), val, flag)
}
func (p *Dao) GetSiteInfo(key, lang string) string {
	var val string
	p.Get(p.site_key(key, lang), &val)
	if val == "" {
		return key
	} else {
		return val
	}
}
func (p *Dao) Set(key string, val interface{}, flag bool) error {
	buf, err := core.ToBits(val)
	if err != nil {
		return err
	}

	s := Setting{Key: key, Flag: flag}
	if flag {
		if v, e := p.Aes.Encrypt(buf); e == nil {
			s.Val = v
		} else {
			return e
		}
	} else {
		s.Val = buf
	}

	var c int
	p.Db.Model(Setting{}).Where("key = ?", key).Count(&c)
	if c == 0 {
		return p.Db.Create(&s).Error
	} else {
		return p.Db.Model(&Setting{}).Where("key = ?", key).UpdateColumns(map[string]interface{}{"val": s.Val}).Error
	}

}

func (p *Dao) Get(key string, val interface{}) error {
	var s Setting
	err := p.Db.Where("key = ?", key).First(&s).Error
	if err != nil {
		return err
	}

	var buf []byte
	if s.Flag {
		if buf, err = p.Aes.Decrypt(s.Val); err != nil {
			return err
		}
	} else {
		buf = s.Val
	}
	return core.FromBits(buf, val)
}

func (*Dao) site_key(key, lang string) string {
	lang = strings.ToLower(lang)
	if lang == "" {
		return fmt.Sprintf("site://%s", key)
	} else {
		return fmt.Sprintf("site://%s/%s", lang, key)
	}
}
