package controllers

import (
    "github.com/astaxie/beego"
    "seraphim/models"
)

type CategoryController struct {
     beego.Controller
}

func (this *CategoryController) Get() {
   this.Data["IsCategory"] = true
   this.TplNames = "category.html"

   var err error
   this.Data["Categories"], err = models.GetAllCategories()
   if err != nil {
      beego.Error(err)
   }

   v := this.GetSession("s1")

   if v == nil {
      this.Data["IsLogin"] = false
      this.Data["uname"] = " "
      return
   }
   if v.(string) != "admin" {
      this.Data["IsLogin"] = true
      this.Data["uname"] = v.(string)
      return
   }

   this.Data["IsLogin"] = true
   this.Data["uname"] = v.(string)

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

}
