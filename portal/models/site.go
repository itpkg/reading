package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Locale struct {
	Id      uint
	Code    string    `orm:"index"`
	Message string    `orm:"type(text)"`
	Lang    string    `orm:"size(5);index"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func (p *Locale) TableUnique() [][]string {
	return [][]string{
		[]string{"Code", "Lang"},
	}
}

type Setting struct {
	Id      uint
	Key     string    `orm:"unique"`
	Val     string    `orm:"type(text)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Locale), new(Setting))
}
