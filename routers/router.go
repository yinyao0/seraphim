package routers

import (
	"seraphim/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})
    beego.Router("/login",&controllers.LoginController{})
    beego.Router("/register",&controllers.RegisterController{})
    beego.Router("/category", &controllers.CategoryController{})
}
