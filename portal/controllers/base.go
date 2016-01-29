package controllers

type BaseController struct {
	LocaleController
}

func (p *BaseController) Prepare() {
	p.LocaleController.Prepare()
	p.Layout = "layout.tpl"
}
