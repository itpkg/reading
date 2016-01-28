package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModelWithPrefix("cms_", new(Article), new(Tag), new(Comment))
}
