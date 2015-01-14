package controllers

import (
     "github.com/astaxie/beego"
     "seraphim/models"
)

type ReplyController struct {
     beego.Controller
}

func (this *ReplyController) Add() {
     tid := this.GetString("tid")
     nickname := this.GetString("nickname")
     content := this.GetString("content")

     err := models.AddReply(tid,nickname,content)
     if err != nil {
       beego.Error(err)
     }
     this.Redirect("/topic/view/"+tid,302)
}

func (this *ReplyController) Delete() {
     v := this.GetSession("s1")
     tid := this.GetString("tid")
     if v == nil {
        this.Redirect("/topic/view/"+tid,302)
        return
     }

     author, err := models.GetAuthorById(tid)
     if err != nil {
        beego.Error(err)
        return
     }
     if v.(string) != author && v.(string) != "admin" {
        this.Redirect("/topic/view/"+tid,302)
        return
     }
     rid := this.GetString("rid")
     err = models.DeleteReply(rid)
     if err != nil {
        beego.Error(err)
     }
     this.Redirect("/topic/view/"+tid,302)
}
