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
    beego.Router("/topic/modify",&controllers.TopicController{},"get:Modify")
    beego.Router("/topic/delete",&controllers.TopicController{},"get:Delete")
    beego.Router("/topic/view/:id:int",&controllers.TopicController{},"get:View")
    beego.Router("/reply/add", &controllers.ReplyController{},"post:Add")
    beego.Router("/reply/delete", &controllers.ReplyController{},"get:Delete")
}
