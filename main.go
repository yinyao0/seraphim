package main

import (
	_ "seraphim/routers"
	"github.com/astaxie/beego"
        "seraphim/models"
)

func init() {
     beego.SessionOn = true
     beego.SessionName = "seraphim"
}

func main() {
        if err := models.InitDB(); err != nil {
           beego.Error(err)
        }
	beego.Run()
}

