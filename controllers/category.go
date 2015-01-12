package controllers

import (
    "github.com/astaxie/beego"
    "seraphim/models"
)

type CategoryController struct {
     beego.Controller
}

func (this *CategoryController) Get() {
   op := this.GetString("op")
   switch op {
     case "add":
          name := this.GetString("name")
          if len(name) == 0 {
             break
          }

         err := models.AddCategory(name)
         if err != nil {
            beego.Error(err)
         }
     this.Redirect("/category", 302)
     return
    case "del":
         id := this.GetString("id")
         if len(id) == 0 {
             break
         }

         err := models.DeleteCategory(id)
         if err != nil {
            beego.Error(err)
         }
       this.Redirect("/category",302)
       return
   }
   this.Data["IsCategory"] = true
   this.TplNames = "category.html"
   v := this.GetSession("s1")
   if v == nil {
      this.Data["IsLogin"] = false
   } else {
      this.Data["IsLogin"] = true
   }

   var err error
   this.Data["Categories"], err = models.GetAllCategories()
   if err != nil {
      beego.Error(err)
   }
}
