package controllers

import (
	"github.com/astaxie/beego"
        "seraphim/models"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
        v := this.GetSession("s1")
        if v != nil {
           this.Redirect("/",302)
           return
        }
	this.TplNames = "register.html"
}

func (this *RegisterController) Post() {
    uname := this.GetString("uname")
    pwd := this.GetString("pwd")
    //this.Ctx.WriteString(uname+pwd)
    var u *models.Users
    u = new(models.Users)
    u = models.GetUserByName(uname)
    if u == nil {
       u = new(models.Users)
       u.Name = uname
       u.Password = pwd
       u.Save()
    } else {
      beego.Debug(u)
      this.Redirect("/register",302)
      return
    }
    beego.Info(u)
    this.Redirect("/",302)
    return
}


