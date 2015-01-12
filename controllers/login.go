package controllers

import (
	"github.com/astaxie/beego"
        "seraphim/models"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
        if this.GetString("exit") == "true" {
           this.DelSession("s1")
           this.Redirect("/",302)
           return
        }
	this.TplNames = "login.html"
}

func (this *LoginController) Post() {
    uname := this.GetString("uname")
    pwd := this.GetString("pwd")
    //this.Ctx.WriteString(uname+pwd)

    if uname == beego.AppConfig.String("adminname") &&
       pwd == beego.AppConfig.String("adminpwd") {
         this.SetSession("s1",uname)
    }
    u := new(models.Users)
    u = models.GetUserByName(uname)
    if u != nil {
       if uname == u.Name && pwd == u.Password {
          this.SetSession("s1",uname)
       }
    }
    this.Redirect("/",302)
    return
}
