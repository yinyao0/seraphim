
package controllers

import (
    "github.com/astaxie/beego"
    "seraphim/models"
    "strings"
)

type TopicController struct {
     beego.Controller
}

func (this *TopicController) Get() {
    this.Data["IsTopics"] = true
    this.TplNames = "topic.html"
    v := this.GetSession("s1")
    if v == nil {
      this.Data["IsLogin"] = false
      this.Data["uname"] = " "
    } else {
      this.Data["IsLogin"] = true
      this.Data["uname"] = v.(string)
    }

    topics, err := models.GetAllTopics("",false)
    if err != nil {
       beego.Error(err)
    }
    this.Data["Topics"] = topics
}

func (this *TopicController) Post() {
     tid := this.GetString("tid")
     title := this.GetString("title")
     category := this.GetString("category")
     content := this.GetString("content")
     author := this.GetSession("s1").(string)

     var err error
     _,fh, err := this.GetFile("attachment")
     if err != nil {
        beego.Error(err)
     }

     var attachment string
     if fh != nil {
        attachment = fh.Filename
        beego.Info(attachment)
        err = this.SaveToFile("attachment","static/attachment/"+attachment)
        if err != nil {
           beego.Error(err)
        }
     }

     if len(tid) == 0 {
       err = models.AddTopic(title,category,content,author,attachment)
     } else {
       writer, err := models.GetAuthorById(tid)
       if err != nil {
          beego.Error(err)
       }

       if writer == author || author == "admin"{
           err = models.ModifyTopic(tid,title,category,content,attachment)
       }
     }

     if err != nil {
        beego.Error(err)
     }
     this.Redirect("/topic",302)
     /*if len(tid) ==0 {
     this.Ctx.WriteString(title+","+category+","+content+","+author)
     }*/
}

func (this *TopicController) Add() {
    this.TplNames = "topic_add.html"
    v := this.GetSession("s1")
    if v == nil {
      this.Data["IsLogin"] = false
      this.Data["uname"] = " "
    } else {
      this.Data["IsLogin"] = true
      this.Data["uname"] = v.(string)
    }

}

func (this *TopicController) Modify() {
   this.TplNames = "topic_modify.html"
   v := this.GetSession("s1")
    if v == nil {
      this.Data["IsLogin"] = false
      this.Data["uname"] = " "
    } else {
      this.Data["IsLogin"] = true
      this.Data["uname"] = v.(string)
    }

   tid := this.GetString("tid")
   topic, err := models.GetTopic(tid)
   if err != nil {
      beego.Error(err)
      this.Redirect("/",302)
      return
   }
   this.Data["Topic"] = topic
   this.Data["Tid"] = tid
}

func (this *TopicController) Delete() {
   v := this.GetSession("s1")
   tid := this.GetString("tid")

   if v == nil {
      this.Redirect("/topic",302)
      return
   } else {
      writer, err := models.GetAuthorById(tid)
      if err !=nil {
         beego.Error(err)
         return
      }
      if writer == v.(string) || v.(string) == "admin" {
         err := models.DeleteTopic(tid)
         if err != nil {
            beego.Error(err)
            return
         }
         this.Redirect("/topic",302)
         return
      } else {
        this.Redirect("/topic",302)
        return
      }
   }
}

func (this *TopicController) View() {
   v := this.GetSession("s1")
   if v == nil {
      this.Data["IsLogin"] = false
      this.Data["uname"] = " "
   } else {
      this.Data["IsLogin"] = true
      this.Data["uname"] = v.(string)
   }
   requrl := this.Ctx.Request.RequestURI
   i := strings.LastIndex(requrl,"/")
   tid := requrl[i+1:]
   topic, err := models.GetTopic(tid)
   if err != nil {
      beego.Error(err)
      return
   }
   this.Data["Topic"] = topic

   replies, err := models.GetAllReplies(tid)
   if err != nil {
      beego.Error(err)
      return
   }
   this.Data["Replies"] =  replies
   this.TplNames = "topic_view.html"
   //this.Ctx.WriteString("view")
}
