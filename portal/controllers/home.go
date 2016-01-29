package controllers

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
	BaseController
}

func (this *MainController) Get() {
	this.TplName = "home.tpl"
}
