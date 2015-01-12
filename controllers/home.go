
package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
        this.Data["IsHome"] = true
	this.TplNames = "home.html"
        v := this.GetSession("s1")
        if v == nil {
          this.Data["IsLogin"] = false
          this.Data["uname"] = " "
        } else {
          this.Data["IsLogin"] = true
          this.Data["uname"] = v.(string)
        }
        //this.TplNames = "home.html"
}
