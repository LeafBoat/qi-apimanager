package routers

import (
	"github.com/astaxie/beego"
	"github.com/qi/apimanager/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
}
