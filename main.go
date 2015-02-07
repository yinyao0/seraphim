package main

import (
	_ "seraphim/routers"
	"github.com/astaxie/beego"
        "seraphim/models"
        "os"
)

func init() {
     beego.SessionOn = true
     beego.SessionName = "seraphim"
     os.Mkdir("static/attachment",os.ModePerm)
}

func main() {
        if err := models.InitDB(); err != nil {
           beego.Error(err)
        }
	beego.Run()
}

