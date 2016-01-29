package controllers

import (
	"github.com/astaxie/beego"
)

type Author struct {
	Name  string
	Email string
}

type Site struct {
	Lang        string
	Title       string
	SubTitle    string
	Keywords    string
	Description string
	Author      Author
	Copyright   string
}

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	//todo
	this.Data["site"] = &Site{
		Lang: "en-US",
	}
	this.Layout = "layout.tpl"
	this.TplName = "home.tpl"
}
