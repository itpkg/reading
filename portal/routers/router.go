package routers

import (
	"github.com/astaxie/beego"
	"github.com/itpkg/reading/portal/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.NotesController{})
}
