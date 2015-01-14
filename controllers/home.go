
package controllers

import (
	"github.com/astaxie/beego"
        "seraphim/models"
        "strings"
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
        cate := this.GetString("cate")
        topics, err := models.GetAllTopics(cate,true)
        if err != nil {
           beego.Error(err)
        }
        //topic.Content = strings.SplitAfterN(topic.Content," ",10)[0]
        for i:=0;i<len(topics);i++ {
          //beego.Warn(len(topics[i].Content))
          if len(topics[i].Content) > 100 {
          topics[i].Content=strings.Join(strings.Split(topics[i].Content," ")[0:10]," ") 
          }
        }
        this.Data["Topics"] = topics

        categories, err := models.GetAllCategories()
        if err != nil {
           beego.Error(err)
        }
        this.Data["Categories"] = categories
}
