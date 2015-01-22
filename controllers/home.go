
package controllers

import (
	"github.com/astaxie/beego"
        "seraphim/models"
        "strings"
        "strconv"
)

type HomeController struct {
	beego.Controller
}


type Page struct {
   Num       int
   Active    bool
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
        page := this.GetString("page")
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
        n := len(topics)/5+1
        pages := make([]Page, n)
        for i:=0;i<n;i++ {
           pages[i].Num = i+1
           pages[i].Active = false
        }
        //this.Data["Pages"] = pages
        t_topics := make([]*models.Topic,0)
        pre := int64(1)
        next := int64(1)
        j := int64(1)
        if page == "" {
           if len(topics) > 5 {
             t_topics = topics[:5]
           } else {
             t_topics = topics[:]
           }
        } else {
           i, _ := strconv.ParseInt(page,10,64)
           j = i
           p := (i - 1) * 5
           switch i {
              case int64(n):
                    t_topics = topics[p:]
                    if pre != 1 {
                       pre = i-1
                    }
                    next = i
              case 1:
                    t_topics = topics[:5]
                    pre = i
                    next = i+1
              default:
                    t_topics = topics[p:p+5]
                    pre = i-1
                    next = i+1
           }
        }
        pages[j-1].Active = true
        this.Data["Pages"] = pages
        this.Data["Pre"] = pre
        this.Data["Next"] = next
        this.Data["Topics"] = t_topics

        categories, err := models.GetAllCategories()
        if err != nil {
           beego.Error(err)
        }
        this.Data["Categories"] = categories
}
