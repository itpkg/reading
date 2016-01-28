package routers

import (
	"github.com/itpkg/reading/portal/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
