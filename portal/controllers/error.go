package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	BaseController
}

func (p *ErrorController) Error401() {
	p.Data["code"] = "errors.e401"
}
func (p *ErrorController) Error403() {
	p.Data["code"] = "errors.e403"
}
func (p *ErrorController) Error404() {
	p.Data["code"] = "errors.e404"
}
func (p *ErrorController) Error500() {
	p.Data["code"] = "errors.e500"
}
func (p *ErrorController) Error503() {
	p.Data["code"] = "errors.e503"
}

func (p *ErrorController) ErrorDb() {
	p.Data["code"] = "db"
}

func (p *ErrorController) Prepare() {
	p.BaseController.Prepare()
	p.TplName = "errors.tpl"
}

func init() {
	beego.ErrorController(&ErrorController{})
}
