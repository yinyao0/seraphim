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
    beego.Router("/topic",&controllers.TopicController{})
    beego.Router("/topic/add",&controllers.TopicController{},"get:Add")
}
