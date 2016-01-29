package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/itpkg/reading/portal/env"
)

type BaseController struct {
	LocaleController
}

func (p *BaseController) Prepare() {
	p.LocaleController.Prepare()

	o := orm.NewOrm()
	site := Site{}
	if err := env.DAO.Get(o, fmt.Sprintf("site.info.%s", p.Lang), &site); err != nil {
		beego.Error(err)
	}

	p.Data["site"] = &site
	p.Layout = "layout.tpl"
}
